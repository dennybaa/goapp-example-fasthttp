package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var developmentEnv *bool

// Application environment configuration parameters
var appConfig struct {
	Environment string
	Backend     string
	FilePath    string
	LogLevel    string
	Port        string
}

// Get enviroment type
func getEnv() string {
	env := "production"
	if developmentEnv != nil {
		if *developmentEnv == true {
			env = "development"
		}
		return env
	}

	// get environment
	developmentEnv = new(bool)
	es := strings.ToLower(os.Getenv("ENVIRONMENT"))

	if es == "development" || es == "devel" {
		*developmentEnv = true
		env = "development"
	} else {
		*developmentEnv = false
	}

	return env
}

// Check if the environment is development
func isDevelopment() bool {
	if developmentEnv != nil {
		return *developmentEnv
	}
	env := getEnv()
	return (env == "development")
}

func initViper() {
	// Defaults
	viper.AutomaticEnv()
	viper.SetDefault("Backend", "logfile")
	viper.SetDefault("Environment", "production")
	viper.SetDefault("LogLevel", "warn")
	viper.SetDefault("Port", "8080")

	// Set the default FilePath based on backend
	filepath := "app.log"
	backend := strings.ToLower(viper.GetString("Backend"))
	if backend == "sqlite" {
		filepath = "data.db"
	} else if backend != "logfile" {
		panic(fmt.Sprintf("%s: %s", "Unsupported backend", backend))
	}
	viper.SetDefault("FilePath", filepath)

	// Load the environment and unmarshal config
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}
}
