package mqtt

import (
	"log"
	"time"

	influxdb "github.com/influxdata/influxdb/client/v2"
)

func (c *MQTTClient) Stored(name string, device string, sensorType string, data int) {
	// Create a new point batch
	bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database:  c.cfg.dbName,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	tags := map[string]string{
		"name":   name,
		"device": device,
		"type":   sensorType,
		"key":    c.cfg.mqttToken,
	}

	fields := map[string]interface{}{
		"data": data,
	}

	pt, err := influxdb.NewPoint("iot", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.db.Write(bp); err != nil {
		log.Fatal(err)
	}
}
