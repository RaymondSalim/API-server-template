package service

import "github.com/Novometrix/web-server-template/server/repository"

/*
	The service contains the business logic.
	It takes request objects, checks for constraints, and returns a response object.
	The service can have multiple dependencies like for example a mail service or another microservice.
	The service interface and it’s implementation are in the same file.
*/

type Services struct {
	FooService
}

func InitService(r *repository.Repositories) *Services {
	return &Services{
		FooService: NewFooService(r.FooRepository),
	}
}
