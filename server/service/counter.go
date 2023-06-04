package service

import (
	"encoding/json"
	constants "github.com/RaymondSalim/API-server-template/server/constants/nsq"
	"github.com/RaymondSalim/API-server-template/server/models"
	"github.com/RaymondSalim/API-server-template/server/repository"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
)

type CounterService interface {
	GetLastCounter(c *gin.Context) (models.Counter, error)
	PublishAddCounter(c *gin.Context) error
	PublishResetCounter(c *gin.Context) error
	AddCounter() error
	ResetCounter() error
}

type counterService struct {
	counterRepository repository.CounterRepository
	producer          *nsq.Producer
}

func NewCounterService(counterRepo repository.CounterRepository, nsqProducer *nsq.Producer) CounterService {
	return &counterService{
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

	return err
}

func (cs counterService) PublishResetCounter(c *gin.Context) error {
	b, _ := json.Marshal("hi")
	err := cs.producer.Publish(constants.Topics[constants.ResetCounter], b)

	return err
}

func (cs counterService) AddCounter() error {
	err := cs.counterRepository.AddCounter()
	return err
}

func (cs counterService) ResetCounter() error {
	err := cs.counterRepository.ResetCounter()
	return err
}
