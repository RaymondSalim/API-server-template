package repository

import (
	"github.com/Novometrix/web-server-template/server/models"
	"gorm.io/gorm"
)

type CounterRepository interface {
	AddCounter() error
	GetLast() (models.Counter, error)
}

type counterRepository struct {
	db *gorm.DB
}

func NewCounterRepository(db *gorm.DB) CounterRepository {
	return counterRepository{db: db}
}

// AddCounter Obviously bad implementation here
func (cr counterRepository) AddCounter() error {
	var ct models.Counter

	ct, _ = cr.GetLast()
	ct.Count += 1
	ct.ID = 0

	cr.db.Create(&ct)

	return nil
}

func (cr counterRepository) GetLast() (models.Counter, error) {
	var ct models.Counter

	db := cr.db.Last(&ct)

	return ct, db.Error
}
