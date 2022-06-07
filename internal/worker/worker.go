package worker

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"

	cnfg "gitlab.ozon.dev/zBlur/homework-3/logistics/config"
	"gitlab.ozon.dev/zBlur/homework-3/logistics/internal/broker/kafka"
	rpstr "gitlab.ozon.dev/zBlur/homework-3/logistics/internal/repository"
	srvc "gitlab.ozon.dev/zBlur/homework-3/logistics/internal/service"
)

type LogisticsWorker struct {
	config              *cnfg.Config
	producer            sarama.SyncProducer
	removeOrderConsumer *RemoveOrderHandler
}

func New(cfg *cnfg.Config, repository rpstr.Repository, service srvc.Service) (*LogisticsWorker, error) {

	brokerConfig := kafka.NewConfig()
	producer, err := kafka.NewSyncProducer(cfg.Kafka.Brokers.String(), brokerConfig)
	if err != nil {
		return nil, err
	}

	worker := &LogisticsWorker{
		config:   cfg,
		producer: producer,
		removeOrderConsumer: &RemoveOrderHandler{
			producer:   producer,
			repository: repository,
			service:    service,
			config:     cfg,
		},
	}

	return worker, nil
}

func (w *LogisticsWorker) StartConsuming(ctx context.Context) error {

	brokerConfig := kafka.NewConfig()
	removeOrder, err := sarama.NewConsumerGroup(
		w.config.Kafka.Brokers.String(),
		fmt.Sprintf("%s%sCG", w.config.Application.Name, w.config.Kafka.IssueOrderTopics.RemoveOrder),
		brokerConfig,
	)
	if err != nil {
		return err
	}

	go func() {
		for {
			err := removeOrder.Consume(ctx, []string{w.config.Kafka.IssueOrderTopics.RemoveOrder}, w.removeOrderConsumer)
			if err != nil {
				log.Printf("%s consumer error: %v", w.config.Kafka.IssueOrderTopics.RemoveOrder, err)
				time.Sleep(time.Second * 5)
			}
		}
	}()
	go func() {
		for err := range removeOrder.Errors() {
			log.Println(err)
		}
	}()

	return nil
}
