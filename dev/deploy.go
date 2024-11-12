package dev

import (
	"path/filepath"

	"github.com/mikerybka/pkg/util"
)

func Deploy(server string) error {
	// Read config
	configDir := filepath.Join(util.HomeDir(), "b2/mikerybka/src/main/data/servers", server)
	panic(configDir)

	// Create server instance

	// Set next-ip

	// Wait 1m

	// Copy files from config dir to /root/config

	// Run /root/config/provision.sh on the server

	// Run docker compose up -d in /root/config

	// Set DNS records with 30s TTL

	// Wait 1m

	// Delete the old server

	// Set ip and delete next-ip
}
