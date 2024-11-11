package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"log"
	"os"
	// "os/exec"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"net/http"
	// "./local_command"
	// "./ssh_command"
	// "./writefile_command"
)

type Task struct {
	Name   string                 `yaml:"name" json:"name"`
	Params map[string]interface{} `yaml:"params" json:"params"`
	Class  string                 `yaml:"class" json:"class"`
}

type Process struct {
	Name  string `yaml:"name" json:"name"`
	Tasks []Task `yaml:"tasks" json:"tasks"`
}

type ProcessesList struct {
	Processes []Process `yaml:"processes" json:"processes"`
}

var logger *log.Logger

func initLogger() {
    // Initialize the logger
    logFile, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        log.Fatal("Error opening log file: ", err)
    }
    logger = log.New(logFile, "CLIApp: ", log.Ldate|log.Ltime|log.Lshortfile)
}


func main() {
	// Initialize the logger
    initLogger()


	var configFile string
	var port int

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A CLI app to run processes",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the CLI app")
		},
	}

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run processes",
		Run: func(cmd *cobra.Command, args []string) {
			processes, err := readConfig(configFile, logger)
			if err != nil{
				logger.Fatalf("Error reading config file: %v", err)
			}
			if len(processes.Processes) == 0 {
				logger.Fatal("No processes found in the configuration file")
			}

			for _, process := range processes.Processes {
				logger.Printf("Running process: %s\n", process.Name)
				for _, task := range process.Tasks {
					switch task.Class {
					case "localCmd":
						runLocalCommand(task, logger)
					case "writefile":
						ExecuteWriteFile(task, logger)
					case "ssh":
						RunSSHCommand(task, logger)
					default:
						logger.Println("Unknown task class")
					}
				}
				logger.Println("Process completed")
			}
		},
	}

	var serverCmd = &cobra.Command{
		Use:  "server",
		Short: "Starts a server to handle process requests",
		Run: func(cmd *cobra.Command, args []string){
			// Define what happens when the command is executed
			http.HandleFunc("/process", handleProcessRequest)
			addr := fmt.Sprintf(":%d", port)
			fmt.Printf("Server listening on port %d\n")
			if err := http.ListenAndServe(addr, nil); err != nil{
				fmt.Printf("Error starting server: %s\n", err)
			}
		},
	}

	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port number for the server")
    rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "Config file for processes")

	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Println(err)
		os.Exit(1)
	}
}

// func readConfig(filename string) []Process {
// 	yamlfile, err := os.ReadFile(filename)
// 	if err != nil {
// 		logger.Fatalf("Error unmarshalling YAML: %v", err)
// 	}

// 	var yamlprocessesList ProcessesList
// 	err = yaml.Unmarshal(yamlfile, &yamlprocessesList)
// 	if err != nil {
// 		logger.Fatalf("Error unmarshalling YAML: %v", err)
// 	}

// 	jsonfile, err := os.ReadFile(filename)
// 	if err != nil{
// 		logger.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	var jsonProcessesList ProcessesList
// 	err = json.Unmarshal(jsonfile, &jsonProcessesList)
// 	if err != nil{
// 		logger.Fatalf("Error unmarshalling JSON: %v", err)
// 	}

// 	return yamlprocessesList.Processes
// }

func readConfig(filename string, logger *log.Logger)(ProcessesList, error){
	var processesList ProcessesList

	file, err := os.ReadFile(filename)
	if err != nil{
		logger.Fatalf("error reading file: %v", err)
	}

	//Determine the file format based on its extension
	ext := filepath.Ext(filename)
	switch ext{
	case ".yaml", ".yml":
		err = yaml.Unmarshal(file, &processesList)
	case ".json":
		err = json.Unmarshal(file, &processesList)
	default:
		logger.Fatalf("unsupported file format: %s", ext)
	}

	if err != nil{
		logger.Fatalf("error unmarshalling data: %v", err)
	}

	return processesList, nil
}

