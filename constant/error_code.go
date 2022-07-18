package constant

import "fmt"

type ErrorPrefix string

type ErrorCode struct {
	StatusCode int
	Prefix     ErrorPrefix
	Number     int
	Message    string
	Detail     string
}

func (e ErrorCode) FullCode() string {
	return fmt.Sprintf("%s-%d", e.Prefix, e.Number)
}

var (
	ET ErrorPrefix = "ET"
	EC ErrorPrefix = "EC"
	EP ErrorPrefix = "EP"

	ET5001 = ErrorCode{500, ET, 5001, "failed to open image file.", "failed to open image file."}
	ET5002 = ErrorCode{500, ET, 5001, "failed to read image file.", "failed to read image file."}
	ET5003 = ErrorCode{500, ET, 5001, "cannot decode image file", "cannot decode image file"}

	EC4001 = ErrorCode{403, EC, 4001, "cannot detect face", "failed to open image file."}
)
