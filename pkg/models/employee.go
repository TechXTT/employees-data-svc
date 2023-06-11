package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primary_key"`
	FirstName  string    `gorm:"not null"`
	LastName   string    `gorm:"not null"`
	Email      string    `gorm:"not null, unique"`
	HireDate   time.Time `gorm:"not null"`
	Salary     float64   `gorm:"not null"`
	PositionID string    `gorm:"not null"`
	Position   Position  `gorm:"foreignKey:PositionID"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	e.ID = uuid

	return
}
