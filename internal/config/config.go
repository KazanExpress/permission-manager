package config

import (
	"log"
	"os"
	"strings"
)

type ClusterConfig struct {
	Name                string
	ControlPlaneAddress string
	Namespace           string
	OverrideNamespace   bool
	CASource            string
}

type BackendConfig struct {
	Port string
}

// Config contains PermissionManager cluster/server configuration
type Config struct {
	Cluster ClusterConfig
	Backend BackendConfig
}

func New() *Config {
	cfg := &Config{
		Cluster: ClusterConfig{
			Name:                os.Getenv("CLUSTER_NAME"),
			ControlPlaneAddress: os.Getenv("CONTROL_PLANE_ADDRESS"),
			Namespace:           os.Getenv("NAMESPACE"),
			CASource:            os.Getenv("CA_SOURCE"),
			OverrideNamespace:   strings.ToLower(os.Getenv("OVERRIDE_NAMESPACE")) == "true",
		},
		Backend: BackendConfig{
			Port: os.Getenv("PORT"),
		},
	}

	if cfg.Backend.Port == "" {
		log.Fatal("PORT env cannot be empty")
	}

	if cfg.Cluster.Name == "" {
		log.Fatal("CLUSTER_NAME env cannot be empty")
	}

	if cfg.Cluster.CASource == "" {
		cfg.Cluster.CASource = "kubeconfig"
	}

	if cfg.Cluster.Namespace == "" {
		log.Fatal("NAMESPACE env cannot be empty")
	}

	if cfg.Cluster.ControlPlaneAddress == "" {
		log.Fatal("CONTROL_PLANE_ADDRESS env cannot be empty")
	}

	return cfg
}
