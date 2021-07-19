package impl

import (
	"excel_tool/common"
	"excel_tool/logging"
	"excel_tool/models"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"mime/multipart"
	"strings"
	"sync"
)

type ExcelServiceImpl struct {
	wg sync.WaitGroup
}

func (e *ExcelServiceImpl) GetInactiveUser(file *multipart.FileHeader, textarea string, columnIndex, exportColumnIndex, sheetIndex int) (string, error) {
	f, err := excelize.OpenFile(file.Filename)
	if err != nil {
		logging.Logger.Error(err)
	}
	rows, err := f.GetRows(f.GetSheetName(sheetIndex))
	if err != nil {
		logging.Logger.Error(err)
	}
	logging.Logger.Debug(len(rows))
	var inactiveUsers []string
	for _, row := range rows[1:] {
		if row[columnIndex] == "未激活" {
			inactiveUsers = append(inactiveUsers, row[exportColumnIndex])
		}
	}
	logging.Logger.Debug(inactiveUsers)
	//解析群成员
	groupUsers := strings.Split(textarea, ";")
	logging.Logger.Debug(groupUsers)
	inactiveUsers = common.Intersect(inactiveUsers, groupUsers)
	logging.Logger.Debug(inactiveUsers)
	var builder strings.Builder
	for _, username := range inactiveUsers {
		builder.WriteString("@")
		builder.WriteString(username)
		builder.WriteString(" ")

	}
	logging.Logger.Debug(builder.String())
	return builder.String(), nil
}

func (e *ExcelServiceImpl) GetExcelData(file *multipart.FileHeader, index int) (*models.ResponseData, error) {
	f, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return nil, err
	}
	// 获取 Sheet1 上所有单元格
	sheetName := f.GetSheetName(index)
	rows, err := f.GetRows(sheetName)
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
		Sheet: &models.Sheet{
			TableHeader: tableHeader,
			TableData:   tableData,
		},
	}, nil
}

func (e *ExcelServiceImpl) GetSheetList(file *multipart.FileHeader) (*models.ResponseData, error) {
	f, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return nil, err
	}
	// 获取 Sheet1 上所有单元格
	sheetList := f.GetSheetList()
	var sheetSlice []*models.SheetList
	for _, s := range sheetList {
		sheetSlice = append(sheetSlice, &models.SheetList{
			SheetIndex: f.GetSheetIndex(s),
			SheetName:  s,
		})
	}
	return &models.ResponseData{
		SheetList: sheetSlice,
	}, nil
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
