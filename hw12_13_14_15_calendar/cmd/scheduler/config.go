package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type Config struct {
	Logger    LoggerConf    `mapstructure:"logger"`
	Storage   StorageConf   `mapstructure:"storage"`
	DB        DBConf        `mapstructure:"db"`
	Scheduler SchedulerConf `mapstructure:"scheduler"`
	Rmq       RMQConf       `mapstructure:"rmq"`
}

type LoggerConf struct {
	Level string `mapstructure:"level"`
	Path  string `mapstructure:"path"`
}

type StorageConf struct {
	Driver         string `mapstructure:"driver"`
	MigrationsPath string `mapstructure:"migrations_path"`
}

type DBConf struct {
	DBHost     string `mapstructure:"host"`
	DBPort     int    `mapstructure:"port"`
	DBName     string `mapstructure:"name"`
	DBUsername string `mapstructure:"username"`
	DBPassword string `mapstructure:"password"`
}

type SchedulerConf struct {
	RunFrequencyInterval   time.Duration `mapstructure:"runFrequencyInterval"`
	TimeForRemoveOldEvents time.Duration `mapstructure:"timeForRemoveOldEvents"`
}

type RMQConf struct {
	URI             string        `mapstructure:"uri"`
	ConsumerTag     string        `mapstructure:"consumerTag"`
	MaxElapsedTime  string        `mapstructure:"maxElapsedTime"`
	InitialInterval string        `mapstructure:"initialInterval"`
	Multiplier      int           `mapstructure:"multiplier"`
	MaxInterval     time.Duration `mapstructure:"maxInterval"`
	Exchange        ExchangeConf
}

type ExchangeConf struct {
	Name       string `mapstructure:"name"`
	Type       string `mapstructure:"type"`
	QueueName  string `mapstructure:"queueName"`
	BindingKey string `mapstructure:"bindingKey"`
}

func NewConfig() *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}

	return &config
}
