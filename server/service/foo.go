package service

import (
	"github.com/Novometrix/web-server-template/server/models"
	"github.com/Novometrix/web-server-template/server/models/request"
	"github.com/Novometrix/web-server-template/server/models/response"
	"github.com/Novometrix/web-server-template/server/repository"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type FooService interface {
	AddFoo(c *gin.Context, fooRequest request.AddFoo) (*response.FooResponse, error)
	GetFoo(c *gin.Context, id int) (*response.FooResponse, error)
	DeleteFoo(c *gin.Context, id int) error
}

type fooService struct {
	fooRepository repository.FooRepository
	producer      *nsq.Producer
}

func NewFooService(fooRepo repository.FooRepository, nsqProducer *nsq.Producer) FooService {
	return fooService{
		fooRepository: fooRepo,
		producer:      nsqProducer,
	}
}

func (fs fooService) AddFoo(c *gin.Context, fooRequest request.AddFoo) (*response.FooResponse, error) {
	var resp response.FooResponse

	newFoo := models.Foo{
		FooName: fooRequest.Name,
	}

	newFoo, err := fs.fooRepository.CreateFoo(c, newFoo)
	if err != nil {
		log.Error("failed to create foo")
		return &resp, err
	}

	return &response.FooResponse{
		Status: "success",
		Foo:    newFoo,
	}, nil
}

func (fs fooService) GetFoo(c *gin.Context, id int) (*response.FooResponse, error) {
	var resp response.FooResponse

	foo, err := fs.fooRepository.GetFoo(c, id)
	if err != nil {
		log.Error("failed to get foo")
		return &resp, err
	}

	return &response.FooResponse{
		Status: "success",
		Foo:    foo,
	}, nil
}

func (fs fooService) DeleteFoo(c *gin.Context, id int) error {
	err := fs.fooRepository.DeleteFoo(c, id)
	if err != nil {
		log.Error("failed to delete foo")
	}

	return err
}
