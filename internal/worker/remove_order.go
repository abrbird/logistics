package worker

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	cnfg "gitlab.ozon.dev/zBlur/homework-3/logistics/config"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/broker/kafka"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/models"
	rpstr "gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
	srvc "gitlab.ozon.dev/zBlur/homework-3/logistics/internal/service"
	"log"
)

type RemoveOrderHandler struct {
	producer   sarama.SyncProducer
	repository rpstr.Repository
	service    srvc.Service
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

		log.Printf("consumer %s: -> %s: %v",
			i.config.Application.Name,
			i.config.Kafka.IssueOrderTopics.RemoveOrder,
			msg.Value,
		)

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
			log.Print("Unmarshall failed: value=%v, err=%v", string(msg.Value), err)
			continue
		}

		ctx := context.Background()
		issuePointRetrieved := i.service.IssuePoint().RetrieveByAddress(
			ctx,
			i.repository.IssuePoint(),
			issueOrderMessage.Address.Id,
		)

		if issuePointRetrieved.Error != nil {
			log.Printf("no sush IssuePoint: %v", err)
			i.RetryRemoveOrder(issueOrderMessage)
			continue
		}

		if !issuePointRetrieved.IssuePoint.IsAvailable {
			log.Printf("IssuePoint is unavailable: %v", err)
			i.RetryRemoveOrder(issueOrderMessage)
			continue
		}

		orderAvailabilityRetrieved := i.service.OrderAvailability().Retrieve(
			ctx,
			i.repository.OrderAvailability(),
			issueOrderMessage.Address.Id,
			issuePointRetrieved.IssuePoint.Id,
		)

		if orderAvailabilityRetrieved.Error != nil {
			log.Printf("no sush Order available on IssuePoint: %v", err)
			i.RetryRemoveOrder(issueOrderMessage)
			continue
		}

		if orderAvailabilityRetrieved.OrderAvailability.Status == models.Issued {
			log.Printf("order is already issued: %v", err)
			continue
		}

		if orderAvailabilityRetrieved.OrderAvailability.Status == models.Moved {
			log.Printf("order is moved: %v", err)
			continue
		}

		if orderAvailabilityRetrieved.OrderAvailability.Status == models.Available {
			orderAvailabilityRetrieved.OrderAvailability.Status = models.Issued

			err = i.service.OrderAvailability().Update(
				ctx,
				i.repository.OrderAvailability(),
				*orderAvailabilityRetrieved.OrderAvailability,
			)
			if err != nil {
				log.Printf("can not update OrderAvailabilty: %v", err)
				i.RetryRemoveOrder(issueOrderMessage)
				continue
			}

			i.SendMarkOrderIssued(issueOrderMessage)
			continue
		}

		log.Printf("It is impossible!: %v", err)
		i.RetryRemoveOrder(issueOrderMessage)

	}
	return nil
}

func (i *RemoveOrderHandler) RetryRemoveOrder(message kafka.IssueOrderMessage) {
	message.Base.SenderServiceName = i.config.Application.Name
	message.Base.Attempt += 1

	if message.Base.Attempt > 3 {
		log.Printf("reached max attempts: %v", message)
		i.SendUndoIssueOrder(message)
		return
	}

	part, offs, kerr, err := kafka.SendMessage(i.producer, i.config.Kafka.IssueOrderTopics.RemoveOrder, message)
	if err != nil {
		log.Printf("can not send message: %v", err)
		return
	}

	if kerr != nil {
		log.Printf("can not send message: %v", kerr)
		return
	}

	log.Printf("consumer %s: %v -> %v", i.config.Application.Name, part, offs)
	return
}

func (i *RemoveOrderHandler) SendUndoIssueOrder(message kafka.IssueOrderMessage) {
	message.Base.SenderServiceName = i.config.Application.Name

	part, offs, kerr, err := kafka.SendMessage(i.producer, i.config.Kafka.IssueOrderTopics.UndoIssueOrder, message)
	if err != nil {
		log.Printf("can not send message: %v", err)
		return
	}

	if kerr != nil {
		log.Printf("can not send message: %v", kerr)
		return
	}

	log.Printf("consumer %s: %v -> %v", i.config.Application.Name, part, offs)
	return
}

func (i *RemoveOrderHandler) SendMarkOrderIssued(message kafka.IssueOrderMessage) {
	message.Base.SenderServiceName = i.config.Application.Name

	part, offs, kerr, err := kafka.SendMessage(i.producer, i.config.Kafka.IssueOrderTopics.MarkOrderIssued, message)
	if err != nil {
		log.Printf("can not send message: %v", err)
		return
	}

	if kerr != nil {
		log.Printf("can not send message: %v", kerr)
		return
	}

	log.Printf("consumer %s: %v -> %v", i.config.Application.Name, part, offs)
	return
}
