package test

import (
	"excel_tool/common"
	"excel_tool/config"
	"excel_tool/logging"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"testing"
)

func init() {
	config.GetConf()
}

var nameSlice []string

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
	md5Encode := common.GetMD5Encode("test.xlsx")
	logging.Logger.Debug(md5Encode)
	members, err := common.GetRedisUtil().SMembers(common.FileDataKey + md5Encode)
	if err != nil {
		logging.Logger.Error(err)
	}
	logging.Logger.Debug(members)
	if len(members) == 0 {
		f, err := excelize.OpenFile("test.xlsx")
		if err != nil {
			logging.Logger.Error(err)
		}
		rows, err := f.GetRows(f.GetSheetName(0))
		if err != nil {
			logging.Logger.Error(err)
		}
		logging.Logger.Debug(len(rows))
		for _, row := range rows[1:] {
			if row[18] == "未激活" {
				nameSlice = append(nameSlice, row[1])
			}
		}
		logging.Logger.Debug(nameSlice)
		logging.Logger.Debug(len(nameSlice))
		common.GetRedisUtil().SAdd(common.FileDataKey+md5Encode, nameSlice)
	}
}

func TestIntersect(t *testing.T) {
	md5Encode := common.GetMD5Encode("test.xlsx")
	logging.Logger.Debug(md5Encode)
	members, err := common.GetRedisUtil().SMembers(common.FileDataKey + md5Encode)
	if err != nil {
		logging.Logger.Error(err)
	}
	logging.Logger.Debug(members)
	var slice = []string{"52Hzの??", "張西偉", "开心，路子怡和路靓怡的爷爷??", "张佳音", "岁月静好", "??",
		"玉霞", "Smith.北羊南虾13559186903", "时尧", "安巴", "马老表", "灿烂阳光", "可乐小仙女????", "金鹏 13755332705",
		"\uE119落叶知秋\uE118", "王丽丽", "王一澳", "杨宗儒", "宿命", "讨厌和心机重的人交往", "苏克谨", "月??", "\uE32E青春永恒\uE32E\uE034", "宁缺勿滥",
		"虞樱花浅", "冰寒绾妤", "含汝", "東升", "墨染傾城", "轩哥", "迷失的自我", "大牛小满", "真爱", "《宇航》", "毛豆豆", ".", "锦宝", "《想念》", "果之恋\uE32718259354767",
		"碗豆", "A 星讯手机维修＋手机抵押??", "破晓之城", "伤心的雪花", "晨雨函babylove", "♂梅儿√", "郑佳轩", "可爱的我", "棒棒糖??", "东山复喜", "好的呀呀呀", "ᝰꫛꫀꪝ.??",
		"我形我塑\uE00D", "大丽", "\uE110莫\uE110", "曾土娇", "背影里的沉默", "悟道", "晨露", "AA华为手机电脑、监控、广告牌制作", "唯美超市",
		"亿湘人五金机电15187139388", "悲寂", "怡子", "白雪", "【卫衣T恤速干】", "123陈星怡", "冬安.", "小兰兰", "婵如", "阳光正好~定制照片书", "温东林", "海艳美甲美睫，15288672365",
		"黄坡粤强陶瓷:罗康超13822555545", "冬爱", "记忆的痕迹", "江水明", "星语星愿", "石春华", "夏爽15295457038", "OLd Summe", "温柔de守候", "李青", "海慧", "红", "宇洪", "董琳彬",
		"A??专业电工13704270288", "尹", "Brian.cui", "文??少", "纤丝美业 （美容美发美体养生）", "荣珠", "张玉红", "王建", "华~七老皮肤管理顾问", "彭启成", "有你们足也??\uE04E",
		"刘杉", "秋（康师傅饮料）", "谢桃业", "A敏", "海森三七15912729002一一邓", "意~意", "吴莲蓬-惠氏旗舰店+13659760936", "L G F Jam", "AA\uE32E唐唐", "雪似黎花"}

	result := common.Intersect(members, slice)
	logging.Logger.Debug(result)
}
