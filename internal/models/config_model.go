package models

type Config struct {
	LogSources logSourcesConfig `yaml:"log_sources"`
}

type logSourcesConfig struct {
	AwsVpcLogs awsVpcLogConfig `yaml:"aws_vpc_logs"`
}

type awsVpcLogConfig struct {
	DefaultUniqueString string `yaml:"default_unique_string"`
	UniqueStringFields  string `yaml:"unique_string_fields"`
}
