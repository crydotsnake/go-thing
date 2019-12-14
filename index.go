package main

import (
    "image"
    "image/png"
    "os"
)

func main() {
    // Erstellt ein leeres Bild mit den maßen 100x200 pixeln
    myImage := image.NewRGBA(image.Rect(0, 0, 100, 200))

    // outputFile is a File type which satisfies Writer interface
    outputFile, err := os.Create("test.png")
    if err != nil {
    	// Wenn ein fehler im programm vorhanden ist, wird er ausgegeben.
    }

    png.Encode(outputFile, myImage)

    // Schließt die Dateien.
    outputFile.Close()
}
