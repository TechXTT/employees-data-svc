package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	ID           uuid.UUID  `gorm:"primary_key"`
	Title        string     `gorm:"not null, unique"`
	DepartmentID string     `gorm:"not null"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
	Employees    []Employee `gorm:"foreignKey:PositionID" json:"-"`
}

type PositionAudit struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primary_key"`
	Action       string    `gorm:"not null"`
	Title        string    `gorm:"not null, unique"`
	DepartmentID uuid.UUID `gorm:"not null"`
}

func (p *Position) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	p.ID = uuid

	return
}
