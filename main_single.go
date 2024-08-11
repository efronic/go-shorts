package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
)

func main_single() {
	// Record the start time
	start := time.Now()

	// Number of frames to generate
	frames := 500

	// Load the input image
	inputImagePath := "image.jpg" // Change this to the path of your input image
	img, err := gg.LoadImage(inputImagePath)
	if err != nil {
		panic(err)
	}

	// Get image dimensions
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	// Create output directory if it doesn't exist
	outputDir := "frames"
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Loop through each frame
	for i := 0; i < frames; i++ {
		// Calculate the x offset based on the frame index
		t := float64(i) / float64(frames)
		xOffset := int(t * float64(imgWidth))

		// Create a new RGBA image
		dc := gg.NewContext(imgWidth, imgHeight)
		dc.SetRGB(1, 1, 1) // Set background color to white
		dc.Clear()

		// Draw the original image with the xOffset
		dc.DrawImage(img, xOffset, 0)

		// Save the frame to a file
		outputImagePath := filepath.Join(outputDir, fmt.Sprintf("frame_%06d.jpg", i))
		file, err := os.Create(outputImagePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = jpeg.Encode(file, dc.Image(), &jpeg.Options{Quality: 80})
		if err != nil {
			panic(err)
		}

		fmt.Printf("Saved frame %d\n", i)
	}

	// Record the end time
	elapsed := time.Since(start)
	fmt.Printf("All frames processed. Total time taken: %s\n", elapsed)
}
