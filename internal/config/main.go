package config

import (
	"fmt"
	"time"

	"github.com/pronuu/roosevelt/internal/constants"
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
	env, _ := utils.GetEnvVar(constants.Environment, "production")

	var portKey string
	if env == "production" {
		portKey = constants.HerokuPort
	} else {
		portKey = constants.RooseveltPort
	}

	port, _ := utils.GetEnvVar(portKey, "7999")
	host, _ := utils.GetEnvVar(constants.RooseveltAddress, "localhost")
	address := fmt.Sprintf("%s:%s", host, port)

	writeTimeoutStr, _ := utils.GetEnvVar(constants.RooseveltWriteTimeout, "15")
	writeTimeout, _ := time.ParseDuration(writeTimeoutStr)

	readTimeoutStr, _ := utils.GetEnvVar(constants.RooseveltReadTimeout, "15")
	readTimeout, _ := time.ParseDuration(readTimeoutStr)

	idleTimeoutStr, _ := utils.GetEnvVar(constants.RooseveltIdleTimeout, "60")
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
