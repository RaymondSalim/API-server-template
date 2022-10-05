package service

import (
	"encoding/json"
	constants "github.com/RaymondSalim/API-server-template/server/constants/nsq"
	"github.com/RaymondSalim/API-server-template/server/models"
	"github.com/RaymondSalim/API-server-template/server/repository"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type CounterService interface {
	GetLastCounter(c *gin.Context) (models.Counter, error)
	PublishAddCounter(c *gin.Context) error
	AddCounter() error
}

type counterService struct {
	counterRepository repository.CounterRepository
	producer          *nsq.Producer
}

func NewCounterService(counterRepo repository.CounterRepository, nsqProducer *nsq.Producer) CounterService {
	return counterService{
		counterRepository: counterRepo,
		producer:          nsqProducer,
	}
}

func (cs counterService) GetLastCounter(c *gin.Context) (models.Counter, error) {
	ct, err := cs.counterRepository.GetLast()

	return ct, err
}

func (cs counterService) PublishAddCounter(c *gin.Context) error {
	b, _ := json.Marshal("hi")
	err := cs.producer.Publish(constants.Topics[constants.VisitCounter], b)
	if err != nil {
		log.Errorf("failed to publish with err: %v", err)
	}

	return nil
}

func (cs counterService) AddCounter() error {
	_ = cs.counterRepository.AddCounter()
	return nil
}
