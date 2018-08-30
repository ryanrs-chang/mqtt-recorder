package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ryanrs-chang/mqtt-recorder/database/models"
)

var database *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=user sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Sensor{})

	db.FirstOrCreate(&models.Sensor{}, models.Sensor{
		Name:       "sensor-weather-01",
		Device:     "HT11",
		SensorType: "weather",
		Key:        "FrYEDnJP18mAcISvbAIm",
	})
	db.FirstOrCreate(&models.Sensor{}, models.Sensor{
		Name:       "sensor-weather-02",
		Device:     "HT11",
		SensorType: "weather",
		Key:        "FrYEDnJP18mAcISvbAIm",
	})

	database = db
	return db
}
