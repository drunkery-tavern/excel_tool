package test

import (
	"excel_tool/config"
	"excel_tool/logging"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"testing"
)

func init() {
	config.GetConf()
}

func TestParseExcel(t *testing.T) {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	var tableHeader []string
	tableHeader = append(tableHeader, rows[0]...)
	logging.Logger.Debug(tableHeader)
	var tableData [][]string
	for _, row := range rows[1:] {
		var rowSlice []string
		rowSlice = append(rowSlice, row...)
		tableData = append(tableData, rowSlice)
	}
	logging.Logger.Debug(tableData)
}

func TestGetExcelData(t *testing.T) {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		logging.Logger.Error(err)
	}
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		logging.Logger.Error(err)
	}
	logging.Logger.Debug(len(rows))
	var nameSlice []string
	for _, row := range rows[1:] {
		if row[18] == "未激活" {
			nameSlice = append(nameSlice, row[1])
		}
	}
	logging.Logger.Debug(nameSlice)
	logging.Logger.Debug(len(nameSlice))
}
