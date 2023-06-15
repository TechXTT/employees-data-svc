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

type EmployeeAudit struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primary_key"`
	Action     string    `gorm:"not null"`
	EmployeeID string    `gorm:"not null"`
	FirstName  string    `gorm:"not null"`
	LastName   string    `gorm:"not null"`
	Email      string    `gorm:"not null, unique"`
	HireDate   time.Time `gorm:"not null"`
	Salary     float64   `gorm:"not null"`
	PositionID string    `gorm:"not null"`
	Position   Position  `gorm:"foreignKey:PositionID"`
}

type EmployeeSalaryHistory struct {
	ID         uuid.UUID `gorm:"primary_key"`
	EmployeeID string    `gorm:"not null"`
	Employee   Employee  `gorm:"foreignKey:EmployeeID"`
	Salary     float64   `gorm:"not null"`
	Effective  time.Time `gorm:"not null"`
	End        time.Time `gorm:"not null"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	e.ID = uuid

	return
}
