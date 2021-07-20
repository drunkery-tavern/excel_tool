package service

import (
	"excel_tool/models"
	"mime/multipart"
)

type IExcelService interface {
	ParseExcel(file *multipart.FileHeader) (*models.ResponseData, error)
	GetSheetList(file *multipart.FileHeader) (*models.ResponseData, error)
	GetExcelData(file *multipart.FileHeader, index int) (*models.ResponseData, error)
	GetInactiveUser(file *multipart.FileHeader, textarea string, columnIndex, exportColumnIndex, sheetIndex int) (*models.ResponseData, error)
}
