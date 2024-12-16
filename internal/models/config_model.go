package models

type Config struct {
	RabbitMq RabbitmqConfig `yaml:"rabbitmq"`
}

type RabbitmqConfig struct {
	User string `yaml:"user"`
	Port uint16 `yaml:"port"`
	Name string `yaml:"queue_name"`
}
