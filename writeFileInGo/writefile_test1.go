package main

import (
    "fmt"
    "os"
    "text/template"
)

type WriteFileTask struct {
    Name       string                 `yaml:"name"`
    Parameters map[string]interface{} `yaml:"parameters"`
    Class      string                 `yaml:"class"`
}

func executeWriteFile(task WriteFileTask, data interface{}) error {
    // Get template content from parameters
    templateContent, ok := task.Parameters["command"].(string)
    if !ok {
        return fmt.Errorf("template parameter is missing or not a string")
    }

    // Create a new template
    tmpl, err := template.New(task.Name).Parse(templateContent)
    if err != nil {
        return err
    }

    // Execute the template
    file, err := os.Create(task.Parameters["output"].(string))
    if err != nil {
        return err
    }
    defer file.Close()

    // Write the output content to the file
    err = tmpl.Execute(file, data)
    if err != nil {
        return err
    }

    return nil
}


func main() {
    // Sample writefile task configuration
    writefileTask := WriteFileTask{
        Name: "greeting_task",
        Parameters: map[string]interface{}{
            "command": "Run: {{.Name}}",
            "output":   "output.txt",
        },
        Class: "writefile",
    }

    // Data to be used in the template
    data := map[string]interface{}{
        "Name": "Write File",
    }

    // Execute the writefile task with dynamic data
    err := executeWriteFile(writefileTask, data)
    if err != nil {
        fmt.Println("Error executing writefile task:", err)
        return
    }

    fmt.Println("Writefile task executed successfully.")
}
