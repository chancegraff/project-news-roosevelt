package config

import (
	"fmt"
	"time"

	"github.com/pronuu/roosevelt/internal/utils"
)

// Config ...
type Config struct {
	Port         string
	Host         string
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

// NewConfig ...
func NewConfig() *Config {
	env, _ := utils.GetEnvVar("ENVIRONMENT", "production")

	var portKey string
	if env == "production" {
		portKey = "PORT"
	} else {
		portKey = "ROOSEVELT_PORT"
	}

	port, _ := utils.GetEnvVar(portKey, "7999")
	host, _ := utils.GetEnvVar("ROOSEVELT_ADDRESS", "localhost")
	address := fmt.Sprintf("%s:%s", host, port)

	writeTimeoutStr, _ := utils.GetEnvVar("ROOSEVELT_WRITE_TIMEOUT", "15")
	writeTimeout, _ := time.ParseDuration(writeTimeoutStr)

	readTimeoutStr, _ := utils.GetEnvVar("ROOSEVELT_READ_TIMEOUT", "15")
	readTimeout, _ := time.ParseDuration(readTimeoutStr)

	idleTimeoutStr, _ := utils.GetEnvVar("ROOSEVELT_IDLE_TIMEOUT", "60")
	idleTimeout, _ := time.ParseDuration(idleTimeoutStr)

	return &Config{
		Port:         port,
		Host:         host,
		Address:      address,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		IdleTimeout:  idleTimeout,
	}
}
