package domain

import "gocv-sample/constant"

type Error struct {
	code constant.ErrorCode
	err  error
}

func NewError(code constant.ErrorCode, err error) Error {
	return Error{code, err}
}

func (e Error) Error() string {
	return e.Error()
}
