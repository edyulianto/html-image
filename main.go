package main

import (
	"log"
	"os/exec"
)

func main() {
	// Define the command to execute wkhtmltoimage
	cmd := exec.Command("wkhtmltoimage", "--height", "860", "--width", "540", "--disable-smart-width", "--quality", "100", "id-card.html", "output.png")

	// Run the command and capture any errors
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to create image: %v", err.Error())
	}

	log.Println("Image generated successfully")
}
