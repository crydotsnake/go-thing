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
    	// Wenn ein fehler im programm vorhanden ist, wird er ausgegeben.
    }

    png.Encode(outputFile, myImage)
    
    
    // Close the files.
    
    outputFile.Close()
}
