package main

import (
	"log"
	"os"
)



// func log_func(filename string){

//     log.Println("This is a log message.")

// 	file, err := os.ReadFile(filename)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	logger := (file, "prefix: ", log.LstdFlags|log.Lshortfile)
// 	logger.Pr
// }


func main() {
    // Log to standard output by default
    log.Println("This is a log message")

    // Log to a specific file
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Use a custom logger with a file output
    logger := log.New(file, "prefix: ", log.LstdFlags|log.Lshortfile)
    logger.Println("This is a custom log message")

}
