package common

import (
	"errors"
)

type ErrorCode uint

const (
	ServerException ErrorCode = iota + 10001
)

const (
	InvalidRequestParams ErrorCode = iota + 20001
)

const (
	UnknownErr ErrorCode = iota + 90001
)

var Error = map[ErrorCode]error{
	UnknownErr:           errors.New("未知错误"),
	InvalidRequestParams: errors.New("参数非法"),
}

func GetMsg(code ErrorCode) string {
	msg, ok := Error[code]
	if ok {
		return msg.Error()
	}

	return Error[UnknownErr].Error()
}
