package main

import (
	"fmt"
	"github.com/EdlinOrg/prominentcolor"
	"image"
	"log"
	"os"
)

func LoadImage(fileInput string) (image.Image, error) {
	f, err := os.Open(fileInput)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}


func main() {
	// step 1: load the image
	img, err := LoadImage("./study_test/example.jpg")
	if err != nil {
		log.Fatal("Failed to load image", err)
	}

	// step 2: Process it
	colours, err := prominentcolor.Kmeans(img)
	if err != nil {
		log.Fatal("Failed to process image", err)
	}

	fmt.Println("Dominant colours")
	for _, colour := range colours {
		fmt.Println("#" + colour.AsString())
	}
}