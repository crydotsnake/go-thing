package main

import (
	"image"
	"image/png"
	"os"
)

func main() {

	// Creates an empty image in size of 100x200

	myImage := image.NewRGBA(image.Rect(0, 0, 100, 200))

	// outputFile is a File type which satisfies Writer interface

	outputFile, err := os.Create("test.png")
	if err != nil {
		// If theres an error,there will be an output.
	}

	png.Encode(outputFile, myImage)

	// Close the files.

	outputFile.Close()
}
