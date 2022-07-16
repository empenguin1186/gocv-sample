package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"log"
	"time"
)

func main() {
	// load classifier file for face detection
	xmlFile := "./data/haarcascade_frontalface_default.xml"
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	if !classifier.Load(xmlFile) {
		log.Fatalf("Error reading cascade file: %v\n", xmlFile)
	}

	start := time.Now().UnixNano()

	// load face image
	filename := "./images/sample.png"
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		log.Fatalf("Error reading image from: %v\n", filename)
	}
	defer img.Close()

	// execute face detection
	rects := classifier.DetectMultiScale(img)

	end := time.Now().UnixNano()

	fmt.Printf("found %d faces\n", len(rects))
	fmt.Printf("start time: %d nano seconds\n", start)
	fmt.Printf("end time: %d nano seconds\n", end)
}
