package excelApi

import (
	"excel_tool/common"
	"excel_tool/handlers/base"
	"excel_tool/logging"
	"excel_tool/service"
	"excel_tool/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
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
	savePath := common.FileSavePath
	_, err = os.Stat(savePath)
	if !os.IsExist(err) {
		if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
			panic(err)
		}
	}
	dst := path.Join(savePath, file.Filename)
	_ = c.SaveUploadedFile(file, dst)
	data, err := ExcelService.ParseExcel(file)
	//data, err := ExcelService.GetSheetList(file)
	if err != nil {
		logging.Logger.Debug("err:", err)
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, data)
}

func (e *ExcelApi) GetExcelData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logging.Logger.Debug(err)
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	indexString := c.PostForm("index")
	index, err := strconv.Atoi(indexString)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	data, err := ExcelService.GetExcelData(file, index)
	logging.Logger.Debug("err:", err)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.GetTableDataFail)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, data)
}

func (e *ExcelApi) GetInactiveUser(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logging.Logger.Debug(err)
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	textarea := c.PostForm("textarea")
	columnValueString := c.PostForm("columnValue")
	columnIndex, err := strconv.Atoi(columnValueString)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	exportColumnValueString := c.PostForm("exportColumnValue")
	exportColumnIndex, err := strconv.Atoi(exportColumnValueString)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	sheetIndexString := c.PostForm("sheetIndex")
	sheetIndex, err := strconv.Atoi(sheetIndexString)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}

	data, err := ExcelService.GetInactiveUser(file, textarea, columnIndex, exportColumnIndex, sheetIndex)
	logging.Logger.Debug(data)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.GetInactiveUserFail)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, data)
}
