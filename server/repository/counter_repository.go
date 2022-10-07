package repository

import (
	"github.com/RaymondSalim/API-server-template/server/models"
	"gorm.io/gorm"
)

type CounterRepository interface {
	AddCounter() error
	GetLast() (models.Counter, error)
	ResetCounter() error
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

	ct, err := cr.GetLast()
	ct.Count += 1
	ct.ID = 0

	err = cr.db.Create(&ct).Error

	return err
}

func (cr counterRepository) GetLast() (models.Counter, error) {
	var ct models.Counter

	db := cr.db.Last(&ct)

	return ct, db.Error
}

func (cr counterRepository) ResetCounter() error {
	var counters []models.Counter
	var err error

	tx := cr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err = tx.Find(&counters).Error

	err = tx.Delete(&counters).Error

	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
	}

	return err
}
