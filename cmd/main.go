package main

import (
	runtime "github.com/aws/aws-lambda-go/lambda"
	"gocv-sample/domain"
	openapi "gocv-sample/presentation"
	"log"
)

func main() {
	classifierFileName := "./data/haarcascade_frontalface_default.xml"
	faceRecognizer, err := domain.NewFaceRecognizer(classifierFileName)
	if err != nil {
		log.Fatalf("cannot create faceRecognizer, err=%v", err)
	}
	recognizeController := openapi.NewRecognizeController(faceRecognizer)
	runtime.Start(recognizeController.PostAuth)
}
