package config

import (
	"log"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	RedisHost     string `env:"REDIS_HOST"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisPort     int    `env:"REDIS_PORT"`
	CacheExpiry   int    `env:"CACHE_EXPIRY"`
	Topic         string `env:"TOPIC"`
}

var (
	configuration Config
)

func GetConfig() Config {
	return configuration
}

func SetConfig() Config {
	t := reflect.TypeOf(configuration)
	v := reflect.ValueOf(&configuration)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := ViperEnvVariable(field.Tag.Get("env"))

		if value != "" {
			switch field.Type.Kind() {
			case reflect.String:
				v.Elem().FieldByName(field.Name).SetString(value)

			case reflect.Int:
				intVal, err := strconv.Atoi(value)
				if err == nil {
					v.Elem().FieldByName(field.Name).SetInt(int64(intVal))
				}
			}
		}
	}

	return configuration
}

func ViperEnvVariable(key string) string {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
