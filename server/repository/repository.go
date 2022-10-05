package repository

import "gorm.io/gorm"

type Repositories struct {
	FooRepository
	CounterRepository
}

func InitRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		FooRepository:     NewFooRepository(db),
		CounterRepository: NewCounterRepository(db),
	}
}
