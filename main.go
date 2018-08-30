package main

import (
	"github.com/jinzhu/gorm"
	"github.com/ryanrs-chang/mqtt-recorder/database/models"

	db "github.com/ryanrs-chang/mqtt-recorder/database"
	"github.com/ryanrs-chang/mqtt-recorder/mqtt"
)

func subscribe(mqtt *mqtt.MQTTClient, mydb *gorm.DB) {
	snesors := &[]models.Sensor{}
	mydb.Find(&snesors)
	for _, sensor := range *snesors {
		mqtt.Subscribe(&sensor)
	}
}

func main() {
	mydb := db.Init()
	c := mqtt.NewBuilder(mqtt.GetOptions())
	c.Connection()
	// c.Run()

	subscribe(c, mydb)

	defer c.Close()
	defer mydb.Close()

	for {
	}
}
