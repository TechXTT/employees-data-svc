package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	Name      string     `gorm:"not null, unique"`
	Positions []Position `gorm:"foreignKey:DepartmentID" json:"-"`
}

func (d *Department) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	d.ID = uuid

	return
}
