package mqtt

import (
	"fmt"
	"log"
	"os"
	"strconv"

	MQTT_Client "github.com/eclipse/paho.mqtt.golang"
	Influxdb "github.com/influxdata/influxdb/client/v2"
	"github.com/ryanrs-chang/mqtt-recorder/database/models"
)

type MQTTClient struct {
	cfg  *Config
	mqtt MQTT_Client.Client
	db   Influxdb.Client
}

func getMainTopic() string {
	return fmt.Sprintf("%s/sensors/+", GetOptions().mqttToken)
}

func NewBuilder(cfg *Config) *MQTTClient {
	return &MQTTClient{cfg: cfg}
}

func (c *MQTTClient) Connection() {
	// MQTT connect Broker
	log.Println("Broker connect..")
	opts := MQTT_Client.NewClientOptions().AddBroker(c.cfg.mqttBroker)
	opts.SetClientID(c.cfg.mqttID)

	c.mqtt = MQTT_Client.NewClient(opts)
	if token := c.mqtt.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("Broker connected.")

	// Connect influxdb
	log.Println("Database connect..")
	c.db, _ = Influxdb.NewHTTPClient(Influxdb.HTTPConfig{
		Addr:     c.cfg.dbHost,
		Username: c.cfg.dbUsername,
		Password: c.cfg.dbPassword,
	})

	log.Println("Database connected.")
}

func (c *MQTTClient) Run() {
	var cbTopic MQTT_Client.MessageHandler = func(client MQTT_Client.Client, msg MQTT_Client.Message) {
		log.Printf("Topic: %s Msg: %s\n", msg.Topic(), msg.Payload())
	}

	topic := getMainTopic()
	log.Printf("Setting sub: %s", topic)
	if token := c.mqtt.Subscribe(topic, 0, cbTopic); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}
}

func (c *MQTTClient) Subscribe(sensor *models.Sensor) {
	topic := fmt.Sprintf("%s/sensors/%s", GetOptions().mqttToken, sensor.Name)

	cb := func(client MQTT_Client.Client, msg MQTT_Client.Message) {
		if msg.Topic() == topic {
			data, _ := strconv.ParseInt(fmt.Sprintf("%s", msg.Payload()), 10, 64)

			log.Println(msg.Topic(), data)
			c.Stored(
				sensor.Name,
				sensor.Device,
				sensor.SensorType,
				int(data),
			)
		}
	}

	if token := c.mqtt.Subscribe(topic, 0, cb); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}

	log.Printf("Subscribe topic:%s, waitting data receive.\n", topic)
}

func (c *MQTTClient) Close() {
	c.mqtt.Disconnect(250)

	if err := c.db.Close(); err != nil {
		log.Fatal(err)
	}
}
