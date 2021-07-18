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
	GetTableDataFail
	GetInactiveUserFail
)

const (
	UnknownErr ErrorCode = iota + 90001
)

var Error = map[ErrorCode]error{
	UnknownErr:           errors.New("未知错误"),
	InvalidRequestParams: errors.New("参数非法"),
	GetTableDataFail:     errors.New("获取表格数据失败"),
	GetInactiveUserFail:  errors.New("数据匹配失败"),
}

func GetMsg(code ErrorCode) string {
	msg, ok := Error[code]
	if ok {
		return msg.Error()
	}

	return Error[UnknownErr].Error()
}
