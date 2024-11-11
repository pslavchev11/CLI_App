package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

var (
	username      string
	host          string
	privateKeyPath string
	command       string
)

var rootCmd = &cobra.Command{
	Use:   "ssh-command",
	Short: "Generate and execute an SSH command",
	Run: func(cmd *cobra.Command, args []string) {
		runSSHCommand()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "SSH username")
	rootCmd.Flags().StringVarP(&host, "host", "H", "", "SSH host")
	rootCmd.Flags().StringVarP(&privateKeyPath, "private-key", "k", "", "Path to private key file")
	rootCmd.Flags().StringVarP(&command, "command", "c", "", "Command to execute on the remote machine")

	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("host")
	rootCmd.MarkFlagRequired("private-key")
	rootCmd.MarkFlagRequired("command")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
}

func runSSHCommand() {
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatalf("Unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		log.Fatalf("Unable to establish SSH connection: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Fatalf("Failed to run command on remote machine: %v", err)
	}

	fmt.Println(string(output))
}
