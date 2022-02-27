package models

import (
	"gorm.io/gorm"
)

type PlayStyle int

const (
	offensive PlayStyle = iota + 1
	defensive
)

type Stats struct {
	Wins  int16
	Loses int16
	Draws int16
}

type Attributes struct {
	Defence float32
	Attack  float32
	Agility float32
}

type Player struct {
	gorm.Model
	Name        string     `json:"name"`
	PlayerStyle PlayStyle  `json:"playerStyle"`
	MainAgent   int        `json:"mainAgent"`
	Attributes  Attributes `json:"attributes"`
}
