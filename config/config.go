package config

import "os"

var AppPort = "8080"

func Load() {
	if port := os.Getenv("PORT"); port != "" {
		AppPort = port
	}
}
