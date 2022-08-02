package domain

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"gocv-sample/constant"
	"gocv.io/x/gocv"
	"io/ioutil"
	"log"
	"mime/multipart"
)

// FaceRecognizer service class for face detection and recognition
type FaceRecognizer struct {
	classifier *gocv.CascadeClassifier
}

// NewFaceRecognizer creates a FaceRecognizer
func NewFaceRecognizer(classifierFileName string) (*FaceRecognizer, error) {

	// load classifier file for face detection
	classifier := gocv.NewCascadeClassifier()
	if !classifier.Load(classifierFileName) {
		return &FaceRecognizer{}, errors.New("cannot load classifier file")
	}

	return &FaceRecognizer{
		classifier: &classifier,
	}, nil
}

func (f *FaceRecognizer) Recognize(storeId string, fileHeader *multipart.FileHeader) error {
	// open image file
	file, err := fileHeader.Open()
	if err != nil {
		log.Printf("failed to open image file. err=%v", err)
		return NewMyError(err, constant.ET5001)
	}

	log.Println("fileHeader.Open() succeeded.")
	defer file.Close()

	// read date from image file
	imgBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("failed to read image file. err=%v", err)
		return NewMyError(err, constant.ET5002)
	}

	// decode face image for face detection
	log.Println("attempting to decode image file...")
	img, err := gocv.IMDecode(imgBytes, gocv.IMReadColor)
	if err != nil {
		log.Printf("cannot decode image file err=%v", err)
		return NewMyError(err, constant.ET5003)
	}
	defer img.Close()

	// execute face detection
	rects := f.classifier.DetectMultiScale(img)

	// see whether face detected
	if len(rects) < 1 {
		log.Printf("failed to detect face from image. err=%d", len(rects))
		return NewMyError(err, constant.EC4001)
	}

	// output face detection result
	for i, e := range rects {
		log.Printf("rectangle(%d) axis -> %s\n", i+1, e.String())
	}

	// output storeId
	log.Printf("storeId: %s", storeId)

	// search image from Amazon Rekognition.
	output, err := f.SearchFacesByImage(imgBytes)
	if err != nil {
		log.Printf("failed to search image from aws rekognition err=%v", err)
		return NewMyError(err, constant.ET5004)
	}
	numOfFacesMatch := len(output.FaceMatches)
	log.Printf("%d faces match", numOfFacesMatch)

	if numOfFacesMatch < 1 {
		return NewMyError(err, constant.EC4002)
	}

	for _, e := range output.FaceMatches {
		log.Printf("imageId: %v, similarity: %v", e.Face.ImageId, e.Similarity)
	}

	// TODO 正常時の戻り値検討
	return nil
}

func (f *FaceRecognizer) SearchFacesByImage(imgBytes []byte) (*rekognition.SearchFacesByImageOutput, error) {
	ctx := context.TODO()

	// https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/gov2/rekognition/DetectFaces
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return &rekognition.SearchFacesByImageOutput{}, err
	}
	cfg.Region = "ap-northeast-1"

	client := rekognition.NewFromConfig(cfg)
	collectionId := "gocv-sample-collection"
	var faceMatchThreshold float32 = 95.000000
	var maxFaces int32 = 5
	input := &rekognition.SearchFacesByImageInput{
		CollectionId:       &collectionId,
		FaceMatchThreshold: &faceMatchThreshold,
		Image: &types.Image{
			Bytes: imgBytes,
		},
		MaxFaces: &maxFaces,
	}

	return client.SearchFacesByImage(ctx, input)
}
