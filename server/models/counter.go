package models

import "gorm.io/gorm"

const counterTableName = "counter-table"

type Counter struct {
	gorm.Model
	Count int
}

func (Counter) TableName() string {
	return counterTableName
}
