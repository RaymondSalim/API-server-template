package repository

import (
	"github.com/Novometrix/web-server-template/server/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

/*
	The repository package contains an interface describing what you can do with it.
	It also contains the repository type that implements the interface.

	The repository takes care of storing models in the database and retrieving them.
	Only domain models go in and out, the repository takes care of the mapping.
*/

type FooRepository interface {
	GetFoo(c *gin.Context, id int) (models.Foo, error)
	CreateFoo(c *gin.Context, foo models.Foo) (models.Foo, error)
	DeleteFoo(c *gin.Context, id int) error
}

type fooRepository struct {
	db *gorm.DB
}

func NewFooRepository(db *gorm.DB) FooRepository {
	return fooRepository{db: db}
}

func (fr fooRepository) GetFoo(c *gin.Context, id int) (models.Foo, error) {
	var foo models.Foo

	res := fr.db.First(&foo, id)
	if res.Error != nil {
		log.Errorf("failed to get foo with id: %d, with err: %v", id, res.Error.Error())
	}

	return models.Foo{}, res.Error
}

func (fr fooRepository) CreateFoo(c *gin.Context, foo models.Foo) (models.Foo, error) {
	res := fr.db.Create(&foo)
	if res.Error != nil {
		log.Errorf("failed to create foo with name: %s, with err: %v", foo.FooName, res.Error.Error())
	}
	return foo, res.Error
}

func (fr fooRepository) DeleteFoo(c *gin.Context, id int) error {
	var foo models.Foo

	res := fr.db.Delete(&foo, id)
	if res.Error != nil {
		log.Errorf("failed to delete foo with id: %d, with err: %v", id, res.Error.Error())
	}

	return res.Error
}
