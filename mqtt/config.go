package mqtt

type Config struct {
	mqttBroker string
	mqttID     string
	mqttToken  string
	dbHost     string
	dbName     string
	dbUsername string
	dbPassword string
}

func GetOptions() *Config {
	return &Config{
		mqttBroker: "tcp://test.mosquitto.org:1883",
		mqttID:     "master",
		mqttToken:  "FrYEDnJP18mAcISvbAIm",

		dbHost:     "http://localhost:8086",
		dbName:     "iot",
		dbUsername: "user",
		dbPassword: "user",
	}
}
