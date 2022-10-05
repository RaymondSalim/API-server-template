package controller

import "github.com/RaymondSalim/API-server-template/server/service"

type Controllers struct {
	FooController
	CounterController

	HealthController
}

func InitController(s *service.Services) *Controllers {
	healthController := new(HealthController)

	return &Controllers{
		FooController:     NewFooController(s.FooService),
		CounterController: NewCounterController(s.CounterService),
		HealthController:  *healthController,
	}
}
