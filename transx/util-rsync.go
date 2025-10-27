package transx

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// executeCommand executes the given command locally or remotely (via SSH).
func executeCommand(commandToExecute string, endpoint EndpointDetails, options *TransferOptions) ([]byte, error) {
	if strings.TrimSpace(commandToExecute) == "" {
		return nil, fmt.Errorf("command to execute cannot be empty")
	}

	if endpoint.IsRemote() { // Check if it's a remote endpoint
		if strings.TrimSpace(endpoint.GetEndpoint()) == "" {
			return nil, fmt.Errorf("HostIP must be provided for remote command execution on endpoint")
		}

		userHost := endpoint.GetEndpoint()
		var username string
		if options != nil && options.RsyncOptions != nil {
			username = options.RsyncOptions.Username
		}
		if strings.TrimSpace(username) != "" {
			userHost = fmt.Sprintf("%s@%s", username, endpoint.GetEndpoint())
		}

		var sshCmdParts []string
		sshCmdParts = append(sshCmdParts, "ssh") // SSH command
		if options != nil && options.RsyncOptions != nil && strings.TrimSpace(options.RsyncOptions.SSHPrivateKeyPath) != "" {
			sshCmdParts = append(sshCmdParts, "-i", options.RsyncOptions.SSHPrivateKeyPath) // Private key
		}
		if endpoint.GetPort() != 0 { // SSH port (if not 0)
			sshCmdParts = append(sshCmdParts, "-p", strconv.Itoa(endpoint.GetPort()))
		}
		if options != nil && options.RsyncOptions != nil && options.RsyncOptions.InsecureSkipHostKeyVerification { // Skip host key verification option
			sshCmdParts = append(sshCmdParts, "-o", "StrictHostKeyChecking=accept-new")
			sshCmdParts = append(sshCmdParts, "-o", "UserKnownHostsFile=/dev/null")
		}

		// Add timeout for SSH connection
		connectTimeout := 30
		if options != nil && options.RsyncOptions != nil && options.RsyncOptions.ConnectTimeout > 0 {
			connectTimeout = options.RsyncOptions.ConnectTimeout
		}
		sshCmdParts = append(sshCmdParts, "-o", fmt.Sprintf("ConnectTimeout=%d", connectTimeout))

		// For remote commands with sudo, we need the -t option to allocate a pseudo-tty
		if strings.Contains(commandToExecute, "sudo") {
			sshCmdParts = append(sshCmdParts, "-t")
		}

		sshCmdParts = append(sshCmdParts, userHost, commandToExecute) // user@host "command_to_execute"

		cmd := exec.Command(sshCmdParts[0], sshCmdParts[1:]...)
		return cmd.CombinedOutput()
	} else {
		// Local execution
		// Use "sh -c" to handle complex shell commands
		cmd := exec.Command("sh", "-c", commandToExecute)
		return cmd.CombinedOutput()
	}
}
