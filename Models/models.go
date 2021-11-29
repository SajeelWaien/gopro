package models

import (
	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	Name      string
	Abilities []Abilities `gorm:"many2many:agent_abilities;"`
	Ult       string
}

type Abilities struct {
	gorm.Model
	Name        string
	Description string
}
