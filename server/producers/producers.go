package producers

import (
	"github.com/RaymondSalim/API-server-template/config"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

func InitProducers(cfg *config.AppConfig) *nsq.Producer {
	nsqCfg := nsq.NewConfig()

	producer, err := nsq.NewProducer(cfg.NSQ.NSQDUrl, nsqCfg)
	if err != nil {
		log.Panic("failed to initialize producer with error: ", err)
	}

	return producer
}
