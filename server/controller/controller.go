package controller

import "github.com/Novometrix/web-server-template/server/service"

type Controllers struct {
	FooController
	HealthController
}

func InitController(s *service.Services) *Controllers {
	healthController := new(HealthController)

	return &Controllers{
		FooController:    NewFooController(s.FooService),
		HealthController: *healthController,
	}
}
