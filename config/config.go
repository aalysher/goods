package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Postgres Postgres `yaml:"postgres"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func LoadConfig() (Config, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		log.Fatalf("Error parsing config file: %s\n", err)
	}

	fmt.Println(cfg)

	return cfg, nil
}
