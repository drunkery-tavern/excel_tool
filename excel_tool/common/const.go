package common

const (
	InsertColM             = "M"
	InsertColN             = "N"
	DefaultSheetIndex      = 0
	IDIndex                = 0
	WorkIndex              = 9
	WorkLinkIndex          = 10
	EmailAuthorizationCode = "bzwosnofbjsgiijb"
	EmailServerAddr        = "smtp.qq.com:25"
	Sender                 = "1519695805@qq.com"
	EmailHost              = "smtp.qq.com"
	EmailSubject           = "服务异常"
)

const (
	FileDataKey = "file_data_"
)

//code status
const (
	SuccessOK = 1000
)

//urls
const (
	ExcelBaseUrl    = "/excel"
	UploadExcel     = "/multi/upload"
	GetExcelData    = "/table"
	GetInactiveUser = "/inactive/user"
	Upload          = "/simple/upload"
	Check           = "/simple/check"
	Merge           = "/simple/merge"
	ScheduleUpload  = "/schedule/upload"
	SystemFiles     = "/system/files"
)

const (
	FileSavePath = "./data/"
)
