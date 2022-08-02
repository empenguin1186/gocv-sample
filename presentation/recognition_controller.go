package openapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"gocv-sample/domain"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
)

func convertToHttpHeader(rh map[string][]string) http.Header {
	headers := http.Header{}
	for h, values := range rh {
		for _, v := range values {
			log.Printf("%s: %s", h, v)
			headers.Add(h, v)
		}
	}
	return headers
}

type RecognizeController struct {
	faceRecognizer *domain.FaceRecognizer
}

func NewRecognizeController(faceRecognizer *domain.FaceRecognizer) *RecognizeController {
	return &RecognizeController{faceRecognizer: faceRecognizer}
}

func (r *RecognizeController) PostAuth(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := convertToHttpHeader(request.MultiValueHeaders)
	_, params, err := mime.ParseMediaType(headers.Get("Content-Type"))
	if err != nil {
		panic(err)
	}
	// リクエストボディをbase64でデコード
	recBody, err := base64.StdEncoding.DecodeString(request.Body)
	if err != nil {
		panic(err)
	}

	// multipart/form-dataをパース
	boundary := params["boundary"]
	br := bytes.NewReader(recBody)
	mr := multipart.NewReader(br, boundary)
	form, err := mr.ReadForm(2 * 1_000_000) // 2MB
	if err != nil {
		panic(err)
	}

	fileHeader := form.File["image"][0]
	param := form.Value["storeId"][0]
	log.Printf("PARAM: %s", param)

	err = r.faceRecognizer.Recognize(param, fileHeader)
	if myError, ok := err.(domain.MyError); ok {
		body := &V1AuthPost500Response{
			Code:        myError.ErrorCode().FullCode(),
			Message:     myError.ErrorCode().Message,
			Description: myError.ErrorCode().Detail,
		}
		byteBody, _ := json.Marshal(body)

		return events.APIGatewayProxyResponse{
			StatusCode: myError.ErrorCode().StatusCode,
			Body:       string(byteBody),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
