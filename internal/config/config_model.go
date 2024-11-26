package config

type Config struct {
	Redis      redisConfig      `yaml:"redis"`
	LogSources logSourcesConfig `yaml:"log_sources"`
}

type redisConfig struct {
	Address string `yaml:"address"`
	Expiry  int    `yaml:"expiry"`
	DB      string `yaml:"db"`
}

type logSourcesConfig struct {
	AwsVpcLogs awsVpcLogConfig `yaml:"aws_vpc_logs"`
}

type awsVpcLogConfig struct {
	DefaultUniqueString string `yaml:"default_unique_string"`
	UniqueStringFields  string `yaml:"unique_string_fields"`
}
