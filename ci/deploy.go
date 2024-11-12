package ci

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mikerybka/pkg/hetzner"
	"github.com/mikerybka/pkg/twilio"
	"github.com/mikerybka/pkg/util"
)

func Deploy(adminPhone, floatingIP string) error {
	// Notify the admin of the deploy
	twilio.SendSMS(adminPhone, "Deployment started.")

	// Get current server ID
	floatingIPs, err := hetzner.ListFloatingIPs()
	if err != nil {
		return err
	}
	oldServerID := 0
	for _, ip := range floatingIPs {
		if ip.IP == floatingIP {
			oldServerID = ip.Server
			break
		}
	}

	// Remove known_hosts file.
	// This is a workaround because sometimes hetzner will give us a VM with the same IP we had in the past.
	sshKnownHosts := filepath.Join(util.HomeDir(), ".ssh/known_hosts")
	os.Remove(sshKnownHosts)

	// Create new server
	fmt.Println("Creating VM")
	vmName := fmt.Sprintf("server%d", time.Now().Unix())
	newVM, err := hetzner.CreateVM(vmName)
	if err != nil {
		return err
	}
	newIP := newVM.PublicNet.IPv4.IP

	// Wait for boot
	fmt.Println("Booting VM")
	time.Sleep(30 * time.Second)

	// Install docker
	fmt.Println("Installing docker")
	out, err := util.ExecRemote("root", newIP, strings.Join([]string{
		// Update the system
		"apt update",
		"apt upgrade -y",

		// Install Prerequisites
		"apt install -y apt-transport-https ca-certificates curl software-properties-common",

		// Add Docker’s Official GPG Key
		"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg",

		// Add Docker Repository
		`echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null`,

		// Install Docker Engine
		"apt update",
		"apt install -y docker-ce docker-ce-cli containerd.io",
		"reboot now",
	}, " && "))
	if err != nil {
		fmt.Println(string(out))
		twilio.SendSMS(adminPhone, fmt.Sprintf("Error installing docker: %s", err))
		return err
	}

	// Wait for reboot
	fmt.Println("Rebooting")
	time.Sleep(20 * time.Second)

	// Copy data
	fmt.Println("Data transfer starting")
	twilio.SendSMS(adminPhone, "Data transfer starting")
	readOnlyStart := time.Now()
	b, err := util.DownloadDir("root", floatingIP, "data")
	if err != nil {
		return err
	}
	fmt.Println("Download complete. Uploading.")
	err = util.UploadDir("root", newIP, b, "/root/")
	if err != nil {
		return err
	}

	// Start services
	fmt.Println("Starting services")
	out, err = util.ExecRemote("root", newIP, "docker compose -f /root/data/src/main/server/compose.yaml up -d")
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	// Assign the floating IP to the new server
	fmt.Println("Starting downtime")
	downtimeStart := time.Now()
	fmt.Println("Assigning floating IP")
	err = hetzner.AssignFloatingIP(floatingIP, newVM.ID)
	if err != nil {
		return err
	}

	// Configure the eth0 device on the new server
	fmt.Println("Adding ip address to the eth0 device")
	out, err = util.ExecRemote("root", newIP, fmt.Sprintf("ip addr add %s dev eth0", floatingIP))
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	fmt.Println("Downtime over. The server was in read-only mode for", time.Since(readOnlyStart), "and was offline for", time.Since(downtimeStart))
	twilio.SendSMS(adminPhone, fmt.Sprintf("Deployment complete. The server was in read-only mode for %s and was totally down for %s. ✌️", time.Since(readOnlyStart), time.Since(downtimeStart)))

	// Delete the old server
	err = hetzner.DeleteServer(oldServerID)
	if err != nil {
		return err
	}

	return nil
}
