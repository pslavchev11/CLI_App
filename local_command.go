package main

import (
	"log"
	//"os"
	"os/exec"
)

// var logger *log.Logger

// func init() {
//     // Initialize the logger
//     logFile, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
//     if err != nil {
//         log.Fatal("Error opening log file: ", err)
//     }
//     logger = log.New(logFile, "CLIApp: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

func runLocalCommand(task Task, logger *log.Logger) {
	command, ok := task.Params["command"].(string)
	if !ok {
		logger.Println("Error: 'command' parameter not found or not a string")
		return
	}

	cmd := exec.Command("powershell", "-c", command)
	output, err := cmd.CombinedOutput()
	logger.Printf("Command output:\n%s\n", output) // Print command output
	if err != nil {
		logger.Printf("Error executing command: %v\n", err)
		return
	}
}