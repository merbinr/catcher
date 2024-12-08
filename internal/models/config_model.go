package models

type Config struct {
	RabbitMQ rabbitmqConfig `yaml:"rabbitmq"`
}

type rabbitmqConfig struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
	Name string `yaml:"queue_name"`
}