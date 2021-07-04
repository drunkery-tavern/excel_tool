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
	}

}
