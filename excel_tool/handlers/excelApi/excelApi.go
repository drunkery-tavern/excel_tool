package excelApi

import (
	"excel_tool/common"
	"excel_tool/handlers/base"
	"excel_tool/logging"
	"excel_tool/models"
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
	data, err := ExcelService.ParseExcel(file.Filename)
	//data, err := ExcelService.GetSheetList(file)
	if err != nil {
		logging.Logger.Debug("err:", err)
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, data)
}

func (e *ExcelApi) GetExcelData(c *gin.Context) {
	_, file, _ := c.Request.FormFile("file")
	indexString := c.PostForm("index")
	index, _ := strconv.Atoi(indexString)
	data, err := ExcelService.GetExcelData(file, index)
	logging.Logger.Debug("err:", err)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.GetTableDataFail)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, data)
}

func (e *ExcelApi) GetInactiveUser(c *gin.Context) {
	filename := c.PostForm("filename")
	textarea := c.PostForm("textarea")
	columnValueString := c.PostForm("columnValue")
	columnIndex, _ := strconv.Atoi(columnValueString)
	exportColumnValueString := c.PostForm("exportColumnValue")
	exportColumnIndex, _ := strconv.Atoi(exportColumnValueString)
	sheetIndexString := c.PostForm("sheetIndex")
	sheetIndex, _ := strconv.Atoi(sheetIndexString)
	data, err := ExcelService.GetInactiveUser(filename, textarea, columnIndex, exportColumnIndex, sheetIndex)
	logging.Logger.Debug(data)
	if err != nil {
		e.RespFailWithDesc(c, http.StatusBadRequest, common.GetInactiveUserFail)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, data)
}

func (e *ExcelApi) SimpleUploaderUpload(c *gin.Context) {
	var chunk models.SimpleUploader
	_, header, err := c.Request.FormFile("file")
	chunk.Filename = c.PostForm("filename")
	chunk.ChunkNumber = c.PostForm("chunkNumber")
	chunk.CurrentChunkSize = c.PostForm("currentChunkSize")
	chunk.Identifier = c.PostForm("identifier")
	chunk.TotalSize = c.PostForm("totalSize")
	chunk.TotalChunks = c.PostForm("totalChunks")
	var chunkDir = "./chunk/" + chunk.Identifier + "/"
	hasDir, _ := common.PathExists(chunkDir)
	if !hasDir {
		if err := common.CreateDir(chunkDir); err != nil {
			logging.Logger.Error("创建目录失败! err", err)
		}
	}
	chunkPath := chunkDir + chunk.Filename + chunk.ChunkNumber
	err = c.SaveUploadedFile(header, chunkPath)
	if err != nil {
		logging.Logger.Error("切片创建失败! err", err)
		e.RespFailWithDesc(c, http.StatusInternalServerError, common.SliceCreationFailed)
		return
	}
	chunk.CurrentChunkPath = chunkPath
	err = ExcelService.SaveChunk(chunk)
	if err != nil {
		logging.Logger.Error("切片创建失败! err", err)
		e.RespFailWithDesc(c, http.StatusInternalServerError, common.SliceCreationFailed)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, gin.H{
		"message": "切片创建成功",
	})
}

func (e *ExcelApi) CheckFileMd5(c *gin.Context) {
	md5 := c.Query("md5")
	err, chunks, isDone := ExcelService.CheckFileMd5(md5)
	if err != nil {
		logging.Logger.Error("md5读取失败! err", err)
		e.RespFailWithDesc(c, http.StatusInternalServerError, common.ReadMd5Failed)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, gin.H{
		"chunks":  chunks,
		"isDone":  isDone,
		"message": "查询成功",
	})
}

func (e *ExcelApi) MergeFileMd5(c *gin.Context) {
	md5 := c.Query("md5")
	fileName := c.Query("fileName")
	logging.Logger.Debugf("md5:%s,filename:%s", md5, fileName)
	err := ExcelService.MergeFileMd5(md5, fileName)
	if err != nil {
		logging.Logger.Error("md5读取失败! err", err)
		e.RespFailWithDesc(c, http.StatusInternalServerError, common.ReadMd5Failed)
		return
	}
	//e.RespSuccess(c, http.StatusOK, common.SuccessOK, gin.H{
	//	"message": "合并成功",
	//})
	data, err := ExcelService.ParseExcel(fileName)
	if err != nil {
		logging.Logger.Error("err:", err)
		e.RespFailWithDesc(c, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	e.RespSuccess(c, http.StatusOK, common.SuccessOK, data)

}
