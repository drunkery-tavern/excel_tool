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
	SliceCreationFailed
	ReadMd5Failed
	MergeExcelFail
	ScheduleSplitFail
	GetAllFilesFail
)

const (
	UnknownErr ErrorCode = iota + 90001
)

var Error = map[ErrorCode]error{
	UnknownErr:           errors.New("未知错误"),
	InvalidRequestParams: errors.New("参数非法"),
	GetTableDataFail:     errors.New("获取表格数据失败"),
	GetInactiveUserFail:  errors.New("数据匹配失败"),
	SliceCreationFailed:  errors.New("创建切片失败"),
	ReadMd5Failed:        errors.New("读取Md5失败"),
	MergeExcelFail:       errors.New("合并Excel失败"),
	ScheduleSplitFail:    errors.New("班期拆分失败"),
	GetAllFilesFail:      errors.New("获取文件列表失败"),
}

func GetMsg(code ErrorCode) string {
	msg, ok := Error[code]
	if ok {
		return msg.Error()
	}

	return Error[UnknownErr].Error()
}
