package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

// https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
func initViper() {
	viper.SetConfigFile(".env")

	viper.SetDefault("CALENDAR_ID", "primary")
	viper.SetDefault("YEARS_AHEAD", "1")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func viperEnvVariableStr(key string) string {
	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion - str")
	}

	fmt.Println("result of viper: " + value)

	return value
}

func viperEnvVariableInt(key string) int {
	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion - str")
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("cannot atoi the value of key: " + key)
	}

	fmt.Println(fmt.Sprint("result of viper: ", value))

	return valueInt
}
