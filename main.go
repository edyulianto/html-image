package main

import (
	"fmt"
	_ "image/png"
	"os"
	"os/exec"
)

// func main() {
// 	// Define the command to execute wkhtmltoimage
// 	cmd := exec.Command("wkhtmltoimage", "--disable-smart-width", "--quality", "100", "--zoom", "1", "id-card.html", "output.png")

// 	// Run the command and capture any errors
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("Failed to create image: %v", err.Error())
// 	}

// 	log.Println("Image generated successfully")

// 	file, err := os.Open("output.png")
// 	if err != nil {
// 		log.Fatalf("Failed to open file: %v", err)
// 	}
// 	defer file.Close()

// 	image, _, err := image.Decode(file)
// 	if err != nil {
// 		log.Fatalf("Failed to encode image: %v", err)
// 	}

// 	newImage := resize.Resize(227, 0, image, resize.NearestNeighbor)

// 	// Encode uses a Writer, use a Buffer if you need the raw []byte
// 	outFile, err := os.Create("resized.jpg")
// 	if err != nil {
// 		log.Fatalf("Failed to create file: %v", err)
// 	}
// 	defer outFile.Close()
// 	err = jpeg.Encode(outFile, newImage, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to encode image: %v", err)
// 	}
// }

func main() {
	// Define the input HTML file and the output PDF file
	inputHTML := "id-card-design-ktp.html"
	outputPDF := "output.pdf"

	// Create the command to execute wkhtmltopdf
	cmd := exec.Command("wkhtmltopdf", inputHTML, outputPDF)

	// Set the command's stdout and stderr to the system's stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running wkhtmltopdf: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("PDF generated successfully")
}
