package consumers

import (
	"fmt"
	"github.com/Novometrix/web-server-template/config"
	constants "github.com/Novometrix/web-server-template/server/constants/nsq"
	"github.com/Novometrix/web-server-template/server/service"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type messageHandler struct {
	*service.Services
	topic   string
	channel string
}

func InitConsumers(cfg *config.AppConfig, service *service.Services) (consumers []*nsq.Consumer) {

	nsqCfg := nsq.NewConfig()

	c1, _ := registerConsumer(cfg, service, constants.Topics[constants.VisitCounter], constants.Channels[constants.Increment], nsqCfg)
	c2, _ := registerConsumer(cfg, service, constants.Topics[constants.ResetCounter], constants.Channels[constants.Decrement], nsqCfg)

	consumers = append(consumers, c1, c2)

	return consumers
}

func registerConsumer(cfg *config.AppConfig, service *service.Services, topic string, channel string, nsqCfg *nsq.Config) (*nsq.Consumer, error) {
	consumer, err := nsq.NewConsumer(topic, channel, nsqCfg)
	if err != nil {
		log.Panicf("failed to initialize consumer with topic: %s, and channel: %s, with error: %v", topic, topic, err)
	}
	consumer.AddHandler(&messageHandler{
		Services: service,
		topic:    topic,
		channel:  channel,
	})
	err = consumer.ConnectToNSQLookupd(cfg.NSQ.NSQLookupdURL)
	if err != nil {
		log.Panicf("failed to connect to nsqlookupd, with error: %v", err)

	}

	return consumer, err
}

func (h messageHandler) HandleMessage(m *nsq.Message) error {
	fmt.Printf("%+v\n", m)
	fmt.Printf("topic: %s, channel: %s\n", h.topic, h.channel)

	if h.topic == constants.Topics[constants.VisitCounter] {
		err := h.Services.CounterService.AddCounter()
		if err != nil {
			log.Errorf("failed to add counter with error: %v", err)
		}
	}
	return nil
}
