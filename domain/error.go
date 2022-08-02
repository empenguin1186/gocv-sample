package domain

import "gocv-sample/constant"

type MyError struct {
	error
	errorCode constant.ErrorCode
}

func NewMyError(error error, errorCode constant.ErrorCode) MyError {
	return MyError{error: error, errorCode: errorCode}
}

func (m MyError) ErrorCode() constant.ErrorCode {
	return m.errorCode
}
