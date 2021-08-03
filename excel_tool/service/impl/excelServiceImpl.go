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

func (e *ExcelServiceImpl) ScheduleSplit(file *multipart.FileHeader) (filename string, err error) {
	tableHeadData := []string{"id", "姓名", "省", "市", "区", "学校名", "年级", "班级", "得分", "作品链接"}
	////"姓名","省","市","区","学校名","年级","班级","得分","作品链接"
	//type ScheduleInfo struct {
	//	ID       string
	//	Name     string
	//	Province string
	//	City     string
	//	Area     string
	//	School   string
	//	Grade    string
	//	Class    string
	//	Score    string
	//	WorkLink string
	//}
	f, err := excelize.OpenFile(common.FileSavePath + file.Filename)
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	rows, err := f.GetRows(f.GetSheetName(common.DefaultSheetIndex))
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	//var scheduleInfoSlice []*ScheduleInfo
	sheetMap := make(map[string]int64)
	for _, row := range rows[1:] {
		if len(row) == 0 {
			break
		}
		if row[18] == "未激活" {
			continue
		}
		//scheduleInfo := &ScheduleInfo{
		//	ID:       row[0],
		//	Name:     row[2],
		//	Province: row[9],
		//	City:     row[10],
		//	Area:     row[11],
		//	School:   row[12],
		//	Grade:    row[13],
		//	Class:    row[14],
		//	Score:    row[23],
		//	WorkLink: row[24],
		//}
		var rowSlice []string
		rowSlice = append(rowSlice, row[0], row[2], row[9], row[10], row[11], row[12], row[13], row[14], row[23], row[24])
		//scheduleInfoSlice = append(scheduleInfoSlice)
		sheetList := f.GetSheetList()
		if common.SliceFind(sheetList, row[20]) {
			//获取插入的行
			_, currentRow := common.GetCurrentRow(sheetMap, row[20])
			//把数据添加到表格
			err = f.SetSheetRow(row[20], fmt.Sprintf("A%d", currentRow), &rowSlice)
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			sheetMap[row[20]] = currentRow + 1
		} else {
			//创建sheet，赋值表头
			sheetIndex := f.NewSheet(row[20])
			style, err := f.NewStyle(
				`{"alignment":{"horizontal":"center","vertical":"center"}}`,
			)
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			err = f.SetColWidth(row[20], "A", "E", 8)
			err = f.SetColWidth(row[20], "F", "F", 40)
			err = f.SetColWidth(row[20], "G", "I", 10)
			err = f.SetColWidth(row[20], "J", "J", 185)
			if err != nil {
				return "", err
			}
			err = f.SetColStyle(row[20], "A:J", style)
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			fillStyle, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#c5deb5"],"pattern":1},"alignment":{"horizontal":"center","vertical":"center"}}`)
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			cellSlice := []string{"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1", "I1", "J1"}
			for _, cell := range cellSlice {
				err = f.SetCellStyle(row[20], cell, cell, fillStyle)
				if err != nil {
					logging.Logger.Error(err)
					return "", err
				}
			}
			err = f.SetSheetRow(f.GetSheetName(sheetIndex), "A1", &tableHeadData)
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			//把数据添加到表格
			err = f.SetSheetRow(f.GetSheetName(sheetIndex), "A2", &rowSlice)
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			sheetMap[f.GetSheetName(sheetIndex)] = 2 + 1
		}
	}
	f.DeleteSheet(f.GetSheetName(common.DefaultSheetIndex))
	err = f.Save()
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	return file.Filename, nil
}

func (e *ExcelServiceImpl) GetTableHeader(files []*multipart.FileHeader) (*models.ResponseData, error) {
	var tableHeader [][]string
	for _, file := range files {
		f, err := excelize.OpenFile(common.FileSavePath + file.Filename)
		if err != nil {
			return nil, err
		}
		// 获取 Sheet1 上所有单元格
		rows, err := f.GetRows(f.GetSheetName(common.DefaultSheetIndex))
		if err != nil {
			return nil, err
		}
		tableHeader = append(tableHeader, rows[0])
	}
	return &models.ResponseData{
		TableHeader: tableHeader,
	}, nil
}

func (e *ExcelServiceImpl) MergeExcel(files []*multipart.FileHeader, model string) (string, error) {
	switch model {
	case "1":
		return e.MergeBDExcel(files)
	case "2":
		return e.MergeWorkExcel(files)
	default:
		return "", errors.New("未知模式")
	}
}

func (e *ExcelServiceImpl) MergeBDExcel(files []*multipart.FileHeader) (string, error) {
	//渠道 区域
	type SourceData struct {
		Channel string
		Area    string
	}
	f, err := excelize.OpenFile(common.FileSavePath + files[1].Filename)
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	rows, err := f.GetRows(f.GetSheetName(common.DefaultSheetIndex))
	if err != nil {
		return "", err
	}
	var sourceDataList []*SourceData
	for _, row := range rows[1:] {
		if len(row) == 0 {
			break
		}
		sourceDataList = append(sourceDataList, &SourceData{
			Channel: row[1],
			Area:    row[0],
		})
	}

	f1, err := excelize.OpenFile(common.FileSavePath + files[0].Filename)
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	dstRows, err := f1.GetRows(f1.GetSheetName(common.DefaultSheetIndex))
	if err != nil {
		return "", err
	}
	for index, row := range dstRows[1:] {
		if len(row) == 0 {
			break
		}
		for _, sourceData := range sourceDataList {
			if strings.EqualFold(row[5], sourceData.Channel) {
				//给大区设值
				err := f1.SetCellValue(f1.GetSheetName(common.DefaultSheetIndex), fmt.Sprintf("E%d", index+2), sourceData.Area)
				if err != nil {
					return "", err
				}
			}
		}
	}
	err = f1.Save()
	if err != nil {
		return "", err
	}
	return files[0].Filename, nil
}

func (e *ExcelServiceImpl) MergeWorkExcel(files []*multipart.FileHeader) (string, error) {
	type WorkInfo struct {
		ID       string
		Score    int
		WorkLink string
	}
	//获取ID对应的作品和链接
	f, err := excelize.OpenFile(common.FileSavePath + files[1].Filename)
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	cols, err := f.Cols(f.GetSheetName(common.DefaultSheetIndex))
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	currentCol := 0
	var ids, scores, workLinks []string
	for cols.Next() {
		if currentCol == 1 {
			col, err := cols.Rows()
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			ids = append(ids, col[1:]...)
		}
		if currentCol == 10 {
			col, err := cols.Rows()
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			scores = append(scores, col[1:]...)
		}

		if currentCol == 5 {
			col, err := cols.Rows()
			if err != nil {
				logging.Logger.Error(err)
				return "", err
			}
			workLinks = append(workLinks, col[1:]...)
		}
		currentCol++
	}
	var workInfoSlice []*WorkInfo
	for index, id := range ids {
		workInfoSlice = append(workInfoSlice, &WorkInfo{
			ID:       id,
			Score:    common.GetScore(scores[index]),
			WorkLink: workLinks[index],
		})
	}
	file, err := excelize.OpenFile(common.FileSavePath + files[0].Filename)
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	sheetList := file.GetSheetList()
	for _, sheet := range sheetList {
		sheetRows, err := file.GetRows(sheet)
		if err != nil {
			logging.Logger.Error(err)
			return "", err
		}
		for rowNum, row := range sheetRows[1:] {
			if len(row) == 0 {
				break
			}
			for _, workInfo := range workInfoSlice {
				if workInfo.ID == row[0] {
					err := file.SetCellValue(sheet, fmt.Sprintf("J%d", rowNum+2), workInfo.WorkLink)
					if err != nil {
						logging.Logger.Error(err)
						return "", err
					}
					err = file.SetCellInt(sheet, fmt.Sprintf("I%d", rowNum+2), workInfo.Score)
					if err != nil {
						logging.Logger.Error(err)
						return "", err
					}
				}
			}
		}
		////删除id列
		//err = file.RemoveCol(sheet, "A")
		//if err != nil {
		//	logging.Logger.Error(err)
		//	return "", err
		//}
	}
	err = file.Save()
	if err != nil {
		logging.Logger.Error(err)
		return "", err
	}
	return files[0].Filename, nil
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
	md5Encode := common.GetMD5Encode(filename)
	logging.Logger.Debug(md5Encode)
	inactiveUsers, err := common.GetRedisUtil().SMembers(common.FileDataKey + md5Encode)
	if err != nil {
		logging.Logger.Error(err)
	}
	if len(inactiveUsers) == 0 {
		f, err := excelize.OpenFile("./finish/" + filename)
		if err != nil {
			logging.Logger.Error(err)
			return nil, err
		}
		rows, err := f.GetRows(f.GetSheetName(sheetIndex))
		if err != nil {
			logging.Logger.Error(err)
			return nil, err
		}
		logging.Logger.Debug(len(rows))
		for _, row := range rows[1:] {
			if len(row) == 0 {
				break
			}
			if row[columnIndex] == "未激活" {
				inactiveUsers = append(inactiveUsers, row[exportColumnIndex])
			}
		}
		common.GetRedisUtil().SAdd(common.FileDataKey+md5Encode, inactiveUsers)
	}
	//解析群成员
	groupUsers := strings.Split(textarea, ";")
	inactiveUsers = common.Intersect(inactiveUsers, groupUsers)
	count := len(inactiveUsers)
	var builder strings.Builder
	for _, username := range inactiveUsers {
		builder.WriteString("@")
		builder.WriteString(username)
		builder.WriteString(" ")
	}
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
		logging.Logger.Error(err)
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
