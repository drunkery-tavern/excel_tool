package impl

import (
	"excel_tool/models"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"mime/multipart"
	"sync"
)

type ExcelServiceImpl struct {
	wg sync.WaitGroup
}

func (e *ExcelServiceImpl) ParseExcel(file *multipart.FileHeader) (*models.ResponseData, error) {
	f, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return nil, err
	}
	// 获取 Sheet1 上所有单元格
	sheetList := f.GetSheetList()
	rows, err := f.GetRows(sheetList[0])
	if err != nil {
		return nil, err
	}
	var tableHeader []string
	tableHeader = append(tableHeader, rows[0]...)
	var tableData [][]string
	for _, row := range rows[1:] {
		var rowSlice []string
		rowSlice = append(rowSlice, row...)
		tableData = append(tableData, rowSlice)
	}
	return &models.ResponseData{
		SheetNameList: sheetList,
		Sheet: &models.Sheet{
			TableHeader: tableHeader,
			TableData:   tableData,
		},
	}, nil
}

func NewExcelServiceImpl() *ExcelServiceImpl {
	return &ExcelServiceImpl{}
}
