package main

import (
	"fmt"
	"github.com/eastata/demux/internal/logger"
	"net"
	"os"
	"strconv"
)

func config() {
	if os.Getenv("SERVER_ADDRESS") == "" {
		os.Setenv("SERVER_ADDRESS", "127.0.0.1")
	} else {
		if net.ParseIP(os.Getenv("SERVER_ADDRESS")) == nil {
			logger.Fatal(fmt.Sprintf("Environment variable SERVER_ADDRESS=%s must contain a valid IP address",
				os.Getenv("SERVER_ADDRESS")))
		}
	}

	if os.Getenv("SERVER_PORT") == "" {
		os.Setenv("SERVER_PORT", "8080")
	} else {
		port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
		if err != nil {
			logger.Fatal(fmt.Sprintf("Environment variable SERVER_PORT=%s must contain a valid port 0-65535",
				os.Getenv("SERVER_PORT")))
		}
		if port < 1 || port > 65535 {
			logger.Fatal(fmt.Sprintf("Environment variable SERVER_PORT=%v must contain a valid port 0-65535", port))
		}
	}
}
