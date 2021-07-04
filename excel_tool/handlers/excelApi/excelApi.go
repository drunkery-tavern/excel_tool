package excelApi

import (
	"excel_tool/common"
	"excel_tool/handlers/base"
	"excel_tool/service"
	"excel_tool/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

var (
	ExcelService service.IExcelService = impl.NewExcelServiceImpl()
)

type ExcelApi struct {
	base.Handler
}

func NewExcelApi() *ExcelApi {
	return &ExcelApi{}
}

func (e *ExcelApi) UploadExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	dst := path.Join("./", file.Filename)
	_ = c.SaveUploadedFile(file, dst)
	data, err := ExcelService.ParseExcel(file)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, data)
}
