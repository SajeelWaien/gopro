package main

import (
	"time"
)

type Agent struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Abilities []string
	Ult       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
