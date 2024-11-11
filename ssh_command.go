// package main

// import (

// 	"log"
// 	"os"

// 	"golang.org/x/crypto/ssh"
// )

// // var logger *log.Logger

// func init() {
// 	// Initialize the logger
// 	logFile, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
// 	if err != nil {
// 		log.Fatal("Error opening log file: ", err)
// 	}
// 	logger = log.New(logFile, "CLIApp: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

// func RunSSHCommand(task Task) {
// 	username, ok := task.Params["username"].(string)
// 	if !ok {
// 		logger.Println("Error: 'username' parameter not found or not a string")
// 		return
// 	}

// 	host, ok := task.Params["host"].(string)
// 	if !ok {
// 		logger.Println("Error: 'host' parameter not found or not a string")
// 		return
// 	}

// 	privateKeyPath, ok := task.Params["private_key"].(string)
// 	if !ok {
// 		logger.Println("Error: 'private_key' parameter not found or not a string")
// 		return
// 	}

// 	command, ok := task.Params["command"].(string)
// 	if !ok {
// 		logger.Println("Error: 'command' parameter not found or not a string")
// 		return
// 	}

// 	privateKey, err := os.ReadFile(privateKeyPath)
// 	if err != nil {
// 		logger.Fatalf("Unable to read private key: %v", err)
// 	}

// 	signer, err := ssh.ParsePrivateKey(privateKey)
// 	if err != nil {
// 		logger.Fatalf("Unable to parse private key: %v", err)
// 	}

// 	config := &ssh.ClientConfig{
// 		User: username,
// 		Auth: []ssh.AuthMethod{
// 			ssh.PublicKeys(signer),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	client, err := ssh.Dial("tcp", host+":22", config)
// 	if err != nil {
// 		logger.Fatalf("Unable to establish SSH connection: %v", err)
// 	}
// 	defer client.Close()

// 	session, err := client.NewSession()
// 	if err != nil {
// 		logger.Fatalf("Failed to create session: %v", err)
// 	}
// 	defer session.Close()

// 	output, err := session.CombinedOutput(command)
// 	if err != nil {
// 		logger.Fatalf("Failed to run command on remote machine: %v", err)
// 	}

// 	logger.Println("Command output:\n", string(output))
// }

//-----------------------------------------------------------------------
// func RunSSHCommand(task Task) {
//     // Extract parameters from the task
//     username, ok := task.Params["username"].(string)
//     if !ok {
//         logger.Println("Error: 'username' parameter not found or not a string")
//         return
//     }

//     host, ok := task.Params["host"].(string)
//     if !ok {
//         logger.Println("Error: 'host' parameter not found or not a string")
//         return
//     }

//     privateKeyPath, ok := task.Params["private_key"].(string)
//     if !ok {
//         logger.Println("Error: 'private_key' parameter not found or not a string")
//         return
//     }

//     command, ok := task.Params["command"].(string)
//     if !ok {
//         logger.Println("Error: 'command' parameter not found or not a string")
//         return
//     }

//     passphrase, ok := task.Params["passphrase"].(string) // Add passphrase parameter
//     if !ok {
//         logger.Println("Error: 'passphrase' parameter not found or not a string")
//         return
//     }

//     // Read the private key file
//     privateKey, err := os.ReadFile(privateKeyPath)
//     if err != nil {
//         logger.Fatalf("Unable to read private key: %v", err)
//     }

//     // Parse the private key with passphrase
    // signer, err := ssh.ParsePrivateKeyWithPassphrase(privateKey, []byte(passphrase))
    // if err != nil {
    //     logger.Fatalf("Unable to parse private key: %v", err)
    // }

    // // Create SSH client configuration
    // config := &ssh.ClientConfig{
    //     User: username,
    //     Auth: []ssh.AuthMethod{
    //         ssh.PublicKeys(signer),
    //     },
    //     HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    // }

    // // Establish SSH connection
    // client, err := ssh.Dial("tcp", host+":22", config)
    // if err != nil {
    //     logger.Fatalf("Unable to establish SSH connection: %v", err)
    // }
    // defer client.Close()

    // // Create SSH session
    // session, err := client.NewSession()
    // if err != nil {
    //     logger.Fatalf("Failed to create session: %v", err)
    // }
    // defer session.Close()

    // // Run command on remote machine
    // output, err := session.CombinedOutput(command)
    // if err != nil {
    //     logger.Fatalf("Failed to run command on remote machine: %v", err)
    // }

	//     // Log command output
	//     logger.Println("Command output:\n", string(output))
	// }

	//-----------------------------------------------------------------------

	// package main

	// import (
	// 	"crypto/sha256"
	// 	"encoding/hex"
	// 	"log"
	// 	"os"
	// 	"golang.org/x/crypto/ssh"
	// 	"fmt"
	// )
	
	// // var logger *log.Logger
	
	// func init() {
	// 	// Initialize the logger
	// 	logFile, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// 	if err != nil {
	// 		log.Fatal("Error opening log file: ", err)
	// 	}
	// 	logger = log.New(logFile, "CLIApp: ", log.Ldate|log.Ltime|log.Lshortfile)
	// }
	
	// func RunSSHCommand(task Task) {
	// 	username, ok := task.Params["username"].(string)
	// 	if !ok {
	// 		logger.Println("Error: 'username' parameter not found or not a string")
	// 		return
	// 	}
	
	// 	host, ok := task.Params["host"].(string)
	// 	if !ok {
	// 		logger.Println("Error: 'host' parameter not found or not a string")
	// 		return
	// 	}
	
	// 	privateKeyPath, ok := task.Params["private_key"].(string)
	// 	if !ok {
	// 		logger.Println("Error: 'private_key' parameter not found or not a string")
	// 		return
	// 	}
	
	// 	command, ok := task.Params["command"].(string)
	// 	if !ok {
	// 		logger.Println("Error: 'command' parameter not found or not a string")
	// 		return
	// 	}
	
	// 	// Prompt the user for passphrase
	// 	fmt.Printf("Enter passphrase for SSH key: ")
	// 	var providedPassphrase string
	// 	fmt.Scanln(&providedPassphrase)
	
	// 	// Hash the provided passphrase
	// 	providedPassphraseHash := hashPassphrase(providedPassphrase)
	
	// 	// Hash the stored passphrase
	// 	storedPassphraseHash, ok := task.Params["passphrase_hash"].(string)
	// 	if !ok {
	// 		logger.Println("Error: 'passphrase_hash' parameter not found or not a string")
	// 		return
	// 	}
	
	// 	// Compare hashes
	// 	if providedPassphraseHash != storedPassphraseHash {
	// 		logger.Println("Error: Incorrect passphrase")
	// 		return
	// 	}
	
	// 	// Read the private key file
	// 	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	// 	if err != nil {
	// 		logger.Fatalf("Unable to read private key: %v", err)
	// 	}
	
	// 	// Parse the private key
	// 	signer, err := ssh.ParsePrivateKey(privateKeyBytes)
	// 	if err != nil {
	// 		logger.Fatalf("Unable to parse private key: %v", err)
	// 	}
	
	// 	// SSH client configuration
	// 	config := &ssh.ClientConfig{
	// 		User: username,
	// 		Auth: []ssh.AuthMethod{
	// 			ssh.PublicKeys(signer),
	// 		},
	// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// 	}
	
	// 	// Connect to the SSH server
	// 	conn, err := ssh.Dial("tcp", host+":22", config)
	// 	if err != nil {
	// 		logger.Fatalf("Failed to dial SSH server: %v", err)
	// 	}
	// 	defer conn.Close()
	
	// 	// Create a session
	// 	session, err := conn.NewSession()
	// 	if err != nil {
	// 		logger.Fatalf("Failed to create SSH session: %v", err)
	// 	}
	// 	defer session.Close()
	
	// 	// Execute the command
	// 	output, err := session.CombinedOutput(command)
	// 	if err != nil {
	// 		logger.Fatalf("Failed to execute command on remote server: %v", err)
	// 	}
	
	// 	// Log command output
	// 	logger.Println("Command output:\n", string(output))
	// }
	
	
	// func hashPassphrase(passphrase string) string {
	// 	hasher := sha256.New()
	// 	hasher.Write([]byte(passphrase))
	// 	return hex.EncodeToString(hasher.Sum(nil))
	// }

	package main

	import (
		"crypto/sha256"
		"encoding/hex"
		"log"
		"os"
		"golang.org/x/crypto/ssh"
		"fmt"
	)
	
	// var logger *log.Logger
	
	// func init() {
	// 	// Initialize the logger
	// 	logFile, err := os.OpenFile("logfile.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// 	if err != nil {
	// 		log.Fatal("Error opening log file: ", err)
	// 	}
	// 	logger = log.New(logFile, "CLIApp: ", log.Ldate|log.Ltime|log.Lshortfile)
	// }
	func RunSSHCommand(task Task, logger *log.Logger) {
		username, ok := task.Params["username"].(string)
		if !ok {
			logger.Println("Error: 'username' parameter not found or not a string")
			return
		}
	
		host, ok := task.Params["host"].(string)
		if !ok {
			logger.Println("Error: 'host' parameter not found or not a string")
			return
		}
	
		privateKeyPath, ok := task.Params["private_key"].(string)
		if !ok {
			logger.Println("Error: 'private_key' parameter not found or not a string")
			return
		}
	
		command, ok := task.Params["command"].(string)
		if !ok {
			logger.Println("Error: 'command' parameter not found or not a string")
			return
		}
	
		// Read the private key file
		privateKeyBytes, err := os.ReadFile(privateKeyPath)
		if err != nil {
			logger.Fatalf("Unable to read private key: %v", err)
		}
	
		// Get the passphrase from user input
		var providedPassphrase string
		fmt.Printf("Enter passphrase for SSH key: ")
		fmt.Scanln(&providedPassphrase)
	
		// Hash the provided passphrase
		providedPassphraseHash := hashPassphrase(providedPassphrase)

		task.Params["passphrase_hash"] = providedPassphraseHash
	
		// Parse the private key
		signer, err := ssh.ParsePrivateKeyWithPassphrase(privateKeyBytes, []byte(providedPassphrase))
		if err != nil {
			logger.Fatalf("Unable to parse private key: %v", err)
		}
	
		// SSH client configuration
		config := &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	
		// Connect to the SSH server
		conn, err := ssh.Dial("tcp", host+":22", config)
		if err != nil {
			logger.Fatalf("Failed to dial SSH server: %v", err)
		}
		defer conn.Close()
	
		// Create a session
		session, err := conn.NewSession()
		if err != nil {
			logger.Fatalf("Failed to create SSH session: %v", err)
		}
		defer session.Close()
	
		// Execute the command
		output, err := session.CombinedOutput(command)
		if err != nil {
			logger.Fatalf("Failed to execute command on remote server: %v", err)
		}
	
		// Log command output
		logger.Println("Command output:\n", string(output))
	}
	func hashPassphrase(passphrase string) string {
		hasher := sha256.New()
		hasher.Write([]byte(passphrase))
		return hex.EncodeToString(hasher.Sum(nil))
	}
	

