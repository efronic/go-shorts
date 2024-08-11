package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/fogleman/gg"
)

func main_multi() {
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

	// Set the number of CPU cores to use
	numWorkers := 4 // Set this to the number of cores you want to use
	runtime.GOMAXPROCS(numWorkers)
	fmt.Printf("Using %d CPU cores\n", numWorkers)

	// Determine the number of available CPU cores / uncomment if you want to max out to all available cores
	// numWorkers := runtime.NumCPU()
	// fmt.Printf("Using %d CPU cores\n", numWorkers)

	// Create a wait group to synchronize the completion of all goroutines
	var wg sync.WaitGroup

	// Create a channel to pass frame indices to the workers
	frameIndices := make(chan int, frames)

	// Start workers
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range frameIndices {
				processFrame(i, img, imgWidth, imgHeight, outputDir, frames)
			}
		}()
	}

	// Send frame indices to the workers
	for i := 0; i < frames; i++ {
		frameIndices <- i
	}
	close(frameIndices)

	// Wait for all workers to finish
	wg.Wait()

	// Record the end time
	elapsed := time.Since(start)
	fmt.Printf("All frames processed. Total time taken: %s\n", elapsed)
}

// processFrame processes a single frame
func processFrame(frameIndex int, img image.Image, imgWidth, imgHeight int, outputDir string, frames int) {
	// Calculate the x offset based on the frame index
	t := float64(frameIndex) / float64(frames)
	xOffset := int(t * float64(imgWidth))

	// Create a new RGBA image
	dc := gg.NewContext(imgWidth, imgHeight)
	dc.SetRGB(1, 1, 1) // Set background color to white
	dc.Clear()

	// Draw the original image with the xOffset
	dc.DrawImage(img, xOffset, 0)

	// Save the frame to a file
	outputImagePath := filepath.Join(outputDir, fmt.Sprintf("frame_%06d.jpg", frameIndex))
	file, err := os.Create(outputImagePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = jpeg.Encode(file, dc.Image(), &jpeg.Options{Quality: 80})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Saved frame %d\n", frameIndex)
}
