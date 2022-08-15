package main

import (
	openapi "gocv-sample/presentation"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	classifierFileName := "./data/haarcascade_frontalface_default.xml"
	DefaultApiService, err := openapi.NewDefaultApiService(classifierFileName)
	if err != nil {
		log.Fatalf("cannot create DefaultApiService, err=%v", err)
	}
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
