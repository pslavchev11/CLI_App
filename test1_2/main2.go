package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Task struct {
	Name   string                 `yaml:"name"`
	Params map[string]interface{} `yaml:"params"`
	Class  string                 `yaml:"class"`
}

type Process struct {
	Name  string `yaml:"name"`
	Tasks []Task `yaml:"tasks"`
}

type ProcessesList struct {
	Processes []Process `yaml:"processes"`
}

var (
	username      string
	host          string
	privateKeyPath string
	command       string
	// configFile    string
)

// var rootCmd = &cobra.Command{
// 	Use:   "app",
// 	Short: "A CLI app to run processes and execute SSH commands",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("Welcome to the CLI app")
// 	},
// }

// var runCmd = &cobra.Command{
// 	Use:   "run",
// 	Short: "Run processes",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		processes := readConfig(configFile)

// 		for _, process := range processes {
// 			fmt.Printf("Running process: %s\n", process.Name)
// 			for _, task := range process.Tasks {
// 				switch task.Class {
// 				case "localCmd":
// 					runLocalCommand(task)
// 				case "writeFile":
// 					data := map[string]interface{}{"key": "value"} // Replace with your data
// 					executeWriteFile(task, data)
// 				case "sshCommand":
// 					runSSHCommand(task)
// 				default:
// 					fmt.Println("Unknown task class")
// 				}
// 			}
// 			fmt.Println("Process completed")
// 		}
// 	},
// }

// func init() {
// 	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "Config file for processes")
// 	rootCmd.AddCommand(runCmd)
// 	//rootCmd.AddCommand(sshCmd)
// }

func main() {

	var configFile string

    var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A CLI app to run processes and execute SSH commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the CLI app")
		},
	}
	
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run processes",
		Run: func(cmd *cobra.Command, args []string) {
			processes := readConfig(configFile)
	
			for _, process := range processes {
				fmt.Printf("Running process: %s\n", process.Name)
				for _, task := range process.Tasks {
					switch task.Class {
					case "localCmd":
						runLocalCommand(task)
					case "writeFile":
						data := map[string]interface{}{"key": "value"} // Replace with your data
						executeWriteFile(task, data)
					case "sshCommand":
						runSSHCommand(task)
					default:
						fmt.Println("Unknown task class")
					}
				}
				fmt.Println("Process completed")
			}
		},
	}

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "Config file for processes")
	rootCmd.AddCommand(runCmd)
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "SSH username")
    rootCmd.Flags().StringVarP(&host, "host", "H", "", "SSH host")
    rootCmd.Flags().StringVarP(&privateKeyPath, "private-key", "k", "", "Path to private key file")
    rootCmd.Flags().StringVarP(&command, "command", "c", "", "Command to execute on the remote machine")

    err := rootCmd.MarkFlagRequired("username")
    if err != nil {
        fmt.Println("Error marking 'username' flag as required:", err)
    }

    err = rootCmd.MarkFlagRequired("host")
    if err != nil {
        fmt.Println("Error marking 'host' flag as required:", err)
    }

    err = rootCmd.MarkFlagRequired("private-key")
    if err != nil {
        fmt.Println("Error marking 'private-key' flag as required:", err)
    }

    err = rootCmd.MarkFlagRequired("command")
    if err != nil {
        fmt.Println("Error marking 'command' flag as required:", err)
    }

	

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readConfig(filename string) []Process {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var processesList ProcessesList
	err = yaml.Unmarshal(file, &processesList)
	if err != nil {
		log.Fatal(err)
	}

	return processesList.Processes
}

func runLocalCommand(task Task) error {
	command, ok := task.Params["command"].(string)
	if !ok {
		return fmt.Errorf("Error: 'command' parameter not found or not a string.")
	}

	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Error executing command: %v\n", err)
	}

	fmt.Printf("Command output:\n%s\n", output)
	return nil
}

func executeWriteFile(task Task, data interface{}) error {
	templateContent, ok := task.Params["command"].(string)
	if !ok {
		return fmt.Errorf("command parameter is missing or not a string")
	}

	tmpl, err := template.New(task.Name).Parse(templateContent)
	if err != nil {
		return err
	}

	file, err := os.Create(task.Params["output"].(string))
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

func runSSHCommand(task Task) {
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
