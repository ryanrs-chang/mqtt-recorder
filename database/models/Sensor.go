package models

import "github.com/jinzhu/gorm"

type Sensor struct {
	gorm.Model
	Key        string
	Name       string
	SensorType string
	Device     string
}
