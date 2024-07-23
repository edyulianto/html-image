package main

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"os/exec"

	"github.com/nfnt/resize"
)

func main() {
	// Define the command to execute wkhtmltoimage
	cmd := exec.Command("wkhtmltoimage", "--disable-smart-width", "--quality", "100", "--zoom", "1", "id-card-design.html", "output.png")

	// Run the command and capture any errors
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to create image: %v", err.Error())
	}

	log.Println("Image generated successfully")

	file, err := os.Open("output.png")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}

	newImage := resize.Resize(227, 0, image, resize.NearestNeighbor)

	// Encode uses a Writer, use a Buffer if you need the raw []byte
	outFile, err := os.Create("resized.jpg")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer outFile.Close()
	err = jpeg.Encode(outFile, newImage, nil)
	if err != nil {
		log.Fatalf("Failed to encode image: %v", err)
	}
}
