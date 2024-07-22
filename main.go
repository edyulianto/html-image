package main

import (
	"log"
	"os/exec"
)

func main() {
	// Define the command to execute wkhtmltoimage
	cmd := exec.Command("wkhtmltoimage", "--width", "856", "--height", "540", "--disable-smart-width", "--quality", "100", "id-card2.html", "output.png")

	// Run the command and capture any errors
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	log.Println("Image generated successfully")
}
