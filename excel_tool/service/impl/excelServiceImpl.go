package impl

import (
	"errors"
	"excel_tool/common"
	"excel_tool/dao"
	"excel_tool/logging"
	"excel_tool/models"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"gorm.io/gorm"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"sync"
)

type ExcelServiceImpl struct {
	wg sync.WaitGroup
}

func (e *ExcelServiceImpl) MergeFileMd5(md5 string, fileName string) error {
	finishDir := "./finish/"
	dir := "./chunk/" + md5
	// 如果文件上传成功 不做后续操作 通知成功即可
	if !errors.Is(dao.Db.First(&models.SimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	// 打开切片文件夹
	rd, err := ioutil.ReadDir(dir)
	_ = os.MkdirAll(finishDir, os.ModePerm)
	// 创建目标文件
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	// 关闭文件
	defer fd.Close()
	// 将切片文件按照顺序写入
	for k := range rd {
		content, _ := ioutil.ReadFile(dir + "/" + fileName + strconv.Itoa(k+1))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
		}
	}

	if err != nil {
		return err
	}
	err = dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		// 删除切片信息
		if err = tx.Delete(&models.SimpleUploader{}, "identifier = ? AND is_done = ?", md5, false).Error; err != nil {
			fmt.Println(err)
			return err
		}
		data := models.SimpleUploader{
			Identifier: md5,
			IsDone:     true,
			FilePath:   finishDir + fileName,
			Filename:   fileName,
		}
		// 添加文件信息
		if err = tx.Create(&data).Error; err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})
	err = os.RemoveAll(dir) // 清除切片
	return err
}

func (e *ExcelServiceImpl) CheckFileMd5(md5 string) (err error, uploads []models.SimpleUploader, isDone bool) {
	err = dao.Db.Find(&uploads, "identifier = ? AND is_done = ?", md5, false).Error
	isDone = errors.Is(dao.Db.First(&models.SimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound)
	return err, uploads, !isDone
}

func (e *ExcelServiceImpl) SaveChunk(chunk models.SimpleUploader) error {
	return dao.Db.Create(chunk).Error
}

func (e *ExcelServiceImpl) GetInactiveUser(filename string, textarea string, columnIndex, exportColumnIndex, sheetIndex int) (*models.ResponseData, error) {
	f, err := excelize.OpenFile("./finish/" + filename)
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
	//logging.Logger.Debug(inactiveUsers)
	//解析群成员
	groupUsers := strings.Split(textarea, ";")
	//logging.Logger.Debug(groupUsers)
	inactiveUsers = common.Intersect(inactiveUsers, groupUsers)
	//logging.Logger.Debug(inactiveUsers)
	count := len(inactiveUsers)
	var builder strings.Builder
	for _, username := range inactiveUsers {
		builder.WriteString("@")
		builder.WriteString(username)
		builder.WriteString(" ")
	}
	//logging.Logger.Debug(builder.String())
	return &models.ResponseData{
		Result: builder.String(),
		Count:  count,
	}, nil
}

func (e *ExcelServiceImpl) GetExcelData(file *multipart.FileHeader, index int) (*models.ResponseData, error) {
	f, err := excelize.OpenFile("./finish/" + file.Filename)
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
	f, err := excelize.OpenFile("./finish/" + file.Filename)
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

func (e *ExcelServiceImpl) ParseExcel(filename string) (*models.ResponseData, error) {
	f, err := excelize.OpenFile("./finish/" + filename)
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
	rows, err := f.GetRows(sheetList[0])
	if err != nil {
		return nil, err
	}
	var tableHeader []string
	tableHeader = append(tableHeader, rows[0]...)
	//var tableData [][]string
	//for _, row := range rows[1:] {
	//	var rowSlice []string
	//	rowSlice = append(rowSlice, row...)
	//	tableData = append(tableData, rowSlice)
	//}
	return &models.ResponseData{
		SheetNameList: sheetList,
		Sheet: &models.Sheet{
			TableHeader: tableHeader,
			//TableData:   tableData,
		},
		SheetList: sheetSlice,
		Count:     0,
	}, nil
}

func NewExcelServiceImpl() *ExcelServiceImpl {
	return &ExcelServiceImpl{}
}
