package models

import (
	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	Name      string      `json:"name"`
	Abilities []Abilities `gorm:"many2many:agent_abilities;" json:"abilities"`
	Ult       string      `json:"ult"`
	Class     string      `json:"class"`
}

type Abilities struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
