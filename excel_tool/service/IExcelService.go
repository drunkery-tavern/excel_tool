package service

import (
	"excel_tool/models"
	"mime/multipart"
)

type IExcelService interface {
	ParseExcel(filename string) (*models.ResponseData, error)
	GetSheetList(file *multipart.FileHeader) (*models.ResponseData, error)
	GetExcelData(file *multipart.FileHeader, index int) (*models.ResponseData, error)
	GetInactiveUser(filename string, textarea string, columnIndex, exportColumnIndex, sheetIndex int) (*models.ResponseData, error)
	SaveChunk(chunk models.SimpleUploader) error
	CheckFileMd5(md5 string) (err error, uploads []models.SimpleUploader, isDone bool)
	MergeFileMd5(md5 string, name string) error
	MergeExcel(files []*multipart.FileHeader, model string) (filename string, err error)
	GetTableHeader(files []*multipart.FileHeader) (*models.ResponseData, error)
}
