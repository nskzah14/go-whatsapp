package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

// Simple example showing how to create a connection and check server version
func main() {
	fmt.Println("Go WhatsApp Library - Simple Example")
	fmt.Println("====================================")
	
	// Create a new WhatsApp connection with a 10-second timeout
	fmt.Println("Creating connection...")
	wac, err := whatsapp.NewConn(10 * time.Second)
	if err != nil {
		log.Fatalf("Error creating connection: %v", err)
	}
	
	// Get client version
	clientVersion := wac.GetClientVersion()
	fmt.Printf("Client version: %d.%d.%d\n", clientVersion[0], clientVersion[1], clientVersion[2])
	
	// Check server version
	fmt.Println("Checking server version...")
	serverVersion, err := whatsapp.CheckCurrentServerVersion()
	if err != nil {
		log.Printf("Warning: Could not check server version: %v", err)
	} else {
		fmt.Printf("Server version: %d.%d.%d\n", serverVersion[0], serverVersion[1], serverVersion[2])
	}
	
	fmt.Println("\nConnection created successfully!")
	fmt.Println("Note: To actually use WhatsApp, you would need to call wac.Login() with a QR code.")
	fmt.Println("See the other examples for complete login and messaging functionality.")
}