package worker

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	cnfg "gitlab.ozon.dev/zBlur/homework-3/logistics/config"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/broker/kafka"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/metrics"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	rpstr "gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
	srvc "gitlab.ozon.dev/zBlur/homework-3/logistics/internal/service"
	"log"
)

type RemoveOrderHandler struct {
	producer   sarama.SyncProducer
	repository rpstr.Repository
	service    srvc.Service
	metrics    metrics.Metrics
	config     *cnfg.Config
}

func (i *RemoveOrderHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (i *RemoveOrderHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (i *RemoveOrderHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		ctx := context.Background()

		if msg.Topic != i.config.Kafka.IssueOrderTopics.RemoveOrder {
			log.Printf(
				"topic names does not match: expected - %s, got %s\n",
				i.config.Kafka.IssueOrderTopics.RemoveOrder,
				msg.Topic,
			)
			continue
		}

		var issueOrderMessage kafka.IssueOrderMessage
		err := json.Unmarshal(msg.Value, &issueOrderMessage)
		if err != nil {
			i.metrics.Error()
			log.Print("Unmarshall failed: value=%v, err=%v", string(msg.Value), err)
			continue
		}

		log.Printf("consumer %s: <- %s: %v",
			i.config.Application.Name,
			i.config.Kafka.IssueOrderTopics.RemoveOrder,
			issueOrderMessage,
		)

		record := i.service.OrderAvailability().RemoveOrder(
			ctx,
			i.repository.OrderAvailability(),
			i.repository.IssuePoint(),
			issueOrderMessage.Order.Id,
			issueOrderMessage.Address.Id,
		)
		if record.Error != nil {
			if errors.Is(record.Error, models.RetryError) {
				err = i.RetryRemoveOrder(issueOrderMessage)
				if err != nil {
					err = i.SendUndoIssueOrder(issueOrderMessage)
					if err != nil {
						i.metrics.KafkaError()
						log.Println(err)
					} else {
						log.Printf(
							"consumer %s: -> %s: %v",
							i.config.Application.Name,
							i.config.Kafka.IssueOrderTopics.UndoIssueOrder,
							issueOrderMessage,
						)
					}
				} else {
					log.Printf(
						"consumer %s: -> %s: %v",
						i.config.Application.Name,
						i.config.Kafka.IssueOrderTopics.RemoveOrder,
						issueOrderMessage,
					)
				}
			} else {
				err = i.SendUndoIssueOrder(issueOrderMessage)
				if err != nil {
					i.metrics.KafkaError()
					log.Println(err)
				} else {
					log.Printf(
						"consumer %s: -> %s: %v",
						i.config.Application.Name,
						i.config.Kafka.IssueOrderTopics.UndoIssueOrder,
						issueOrderMessage,
					)
				}
			}
			continue
		}

		err = i.SendMarkOrderIssued(issueOrderMessage)
		if err != nil {
			i.metrics.KafkaError()
			log.Println(err)
		} else {
			log.Printf(
				"consumer %s: -> %s: %v",
				i.config.Application.Name,
				i.config.Kafka.IssueOrderTopics.MarkOrderIssued,
				issueOrderMessage,
			)
		}
	}
	return nil
}

func (i *RemoveOrderHandler) RetryRemoveOrder(message kafka.IssueOrderMessage) error {
	message.Base.SenderServiceName = i.config.Application.Name
	message.Base.Attempt += 1

	if message.Base.Attempt > 5 {
		return models.NewMaxAttemptsError(nil)
	}

	part, offs, kerr, err := kafka.SendMessage(i.producer, i.config.Kafka.IssueOrderTopics.RemoveOrder, message)
	if err != nil {
		return models.BrokerSendError(err)
	}

	if kerr != nil {
		return models.BrokerSendError(err)
	}

	_ = part
	_ = offs

	return nil
}

func (i *RemoveOrderHandler) SendUndoIssueOrder(message kafka.IssueOrderMessage) error {
	message.Base.SenderServiceName = i.config.Application.Name

	part, offs, kerr, err := kafka.SendMessage(i.producer, i.config.Kafka.IssueOrderTopics.UndoIssueOrder, message)
	if err != nil {
		return models.BrokerSendError(err)
	}

	if kerr != nil {
		return models.BrokerSendError(err)
	}
	_ = part
	_ = offs
	return nil
}

func (i *RemoveOrderHandler) SendMarkOrderIssued(message kafka.IssueOrderMessage) error {
	message.Base.SenderServiceName = i.config.Application.Name

	part, offs, kerr, err := kafka.SendMessage(i.producer, i.config.Kafka.IssueOrderTopics.MarkOrderIssued, message)
	if err != nil {
		return models.BrokerSendError(err)
	}

	if kerr != nil {
		return models.BrokerSendError(err)
	}
	_ = part
	_ = offs

	return nil
}
