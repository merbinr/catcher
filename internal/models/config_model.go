package models

type Config struct {
	RabbitMQ   rabbitmqConfig   `yaml:"rabbitmq"`
	LogSources logSourcesConfig `yaml:"log_sources"`
}

type rabbitmqConfig struct {
	User string `yaml:"user"`
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
	Name string `yaml:"queue_name"`
}

type logSourcesConfig struct {
	AwsVpcLogs awsVpcLogConfig `yaml:"aws_vpc_logs"`
}

type awsVpcLogConfig struct {
	DefaultUniqueString string `yaml:"default_unique_string"`
	UniqueStringFields  string `yaml:"unique_string_fields"`
}
