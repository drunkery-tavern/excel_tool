package routers

import (
	"excel_tool/common"
	"excel_tool/handlers/excelApi"
	"github.com/gin-gonic/gin"
)

var (
	ExcelApi = excelApi.NewExcelApi()
)

func excelRouters(r *gin.Engine) {
	excel := r.Group(common.ExcelBaseUrl)
	{
		excel.POST(common.UploadExcel, ExcelApi.UploadExcel)
		excel.POST(common.GetExcelData, ExcelApi.GetExcelData)
		excel.POST(common.GetInactiveUser, ExcelApi.GetInactiveUser)

		excel.POST(common.Upload, ExcelApi.SimpleUploaderUpload)
		excel.GET(common.Check, ExcelApi.CheckFileMd5)
		excel.GET(common.Merge, ExcelApi.MergeFileMd5)

		excel.POST(common.ScheduleUpload, ExcelApi.ScheduleUpload)
		excel.GET(common.SystemFiles, ExcelApi.GetSystemFiles)
	}
}
