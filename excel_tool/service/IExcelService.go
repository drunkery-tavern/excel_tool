package service

import (
	"excel_tool/models"
	"mime/multipart"
)

type IExcelService interface {
	ParseExcel(file *multipart.FileHeader) (*models.ResponseData, error)
}
