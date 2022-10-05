package service

import (
	"github.com/RaymondSalim/API-server-template/server/repository"
	"github.com/nsqio/go-nsq"
)

/*
	The service contains the business logic.
	It takes request objects, checks for constraints, and returns a response object.
	The service can have multiple dependencies like for example a mail service or another microservice.
	The service interface and it’s implementation are in the same file.
*/

type Services struct {
	FooService
	CounterService
}

func InitService(r *repository.Repositories, nsqProducer *nsq.Producer) *Services {
	return &Services{
		FooService:     NewFooService(r.FooRepository, nsqProducer),
		CounterService: NewCounterService(r.CounterRepository, nsqProducer),
	}
}
