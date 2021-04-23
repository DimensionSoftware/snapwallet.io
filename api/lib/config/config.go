package config

import (
	"fmt"
	"log"
	"os"
)

const apiHostEnvVarName = "API_HOST"
const webHostEnvVarName = "WEB_HOST"

type APIHost string
type WebHost string

// ProvideAPIHost ...
func ProvideAPIHost() (APIHost, error) {
	apiHost := os.Getenv(apiHostEnvVarName)
	if apiHost == "" {
		return "", fmt.Errorf("you must set %s", apiHostEnvVarName)
	}

	log.Println("🚨 API Host for webhooks set to: ", apiHost)

	return APIHost(apiHost), nil
}

// ProvideWebHost ...
func ProvideWebHost() (WebHost, error) {
	webHost := os.Getenv(webHostEnvVarName)
	if webHost == "" {
		return "", fmt.Errorf("you must set %s", webHostEnvVarName)
	}

	log.Println("🚨 Web Host for short urls set to: ", webHost)

	return WebHost(webHost), nil
}
