package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

//var logger *log.Logger

// func init() {
//     // Initialize the logger
//     logFile, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
//     if err != nil {
//         log.Fatal("Error opening log file: ", err)
//     }
//     logger = log.New(logFile, "CLIApp: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

func ExecuteWriteFile(task Task, logger *log.Logger) error {
    // Get template content from parameters
    templateContent, ok := task.Params["command"].(string)
    if !ok {
        return fmt.Errorf("template parameter is missing or not a string")
    }

    // Create a new template
    tmpl, err := template.New(task.Name).Parse(templateContent)
    if err != nil {
        return err
    }

    // Execute the template
    file, err := os.Create(task.Params["output"].(string))
    if err != nil {
        return err
    }
    defer file.Close()

    // Write the output content to the file
    err = tmpl.Execute(file, nil)
    if err != nil {
        return err
    }

    logger.Println("File written successfully")
    return nil
}
