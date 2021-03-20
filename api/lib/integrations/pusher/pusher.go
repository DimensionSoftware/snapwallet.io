package pusher

import (
	"fmt"
	"os"

	"github.com/pusher/pusher-http-go"
)

const pusherAppIdEnvVarName = "PUSHER_APP_ID"
const pusherKeyEnvVarName = "PUSHER_KEY"
const pusherSecretEnvVarName = "PUSHER_SECRET"
const pusherClusterEnvVarName = "PUSHER_CLUSTER"

type PusherConfig struct {
	AppID   string
	Key     string
	Secret  string
	Cluster string
}

// ProviderPusherConfig ...
func ProviderPusherConfig() (*PusherConfig, error) {
	pusherAppID, err := requireEnv(pusherAppIdEnvVarName)
	if err != nil {
		return nil, err
	}

	pusherKey, err := requireEnv(pusherKeyEnvVarName)
	if err != nil {
		return nil, err
	}

	pusherSecret, err := requireEnv(pusherSecretEnvVarName)
	if err != nil {
		return nil, err
	}

	pusherCluster, err := requireEnv(pusherClusterEnvVarName)
	if err != nil {
		return nil, err
	}

	return &PusherConfig{
		AppID:   pusherAppID,
		Key:     pusherKey,
		Secret:  pusherSecret,
		Cluster: pusherCluster,
	}, nil
}

// ProvidePusherClient ...
func ProvidePusherClient(config *PusherConfig) *pusher.Client {

	return &pusher.Client{
		AppID:   config.AppID,
		Key:     config.Key,
		Secret:  config.Secret,
		Cluster: config.Cluster,
		Secure:  true,
	}

}

func requireEnv(envVarName string) (string, error) {
	val := os.Getenv(envVarName)
	if val == "" {
		return "", fmt.Errorf("you must set %s", envVarName)
	}

	return val, nil
}
