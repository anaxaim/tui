package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Redis  RedisConfig  `yaml:"redis"`
}

type ServerConfig struct {
	ENV                    string `yaml:"env"`
	Address                string `yaml:"address"`
	Port                   int    `yaml:"port"`
	GracefulShutdownPeriod int    `yaml:"gracefulShutdownPeriod"`
	JWTSecret              string `yaml:"jwtSecret"`
}

type DBConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	Database       string `yaml:"database"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	MigrationsPath string `yaml:"migrationsPath"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (rc RedisConfig) String() string {
	if rc.Host == "" && rc.Port == "" {
		return ""
	}

	return rc.Host + ":" + rc.Port
}

func Parse(appConfig string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(appConfig)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
