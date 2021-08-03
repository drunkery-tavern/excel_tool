package common

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"excel_tool/logging"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func GetRandomString() string {
	n := 6
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	s = strings.ToLower(s)
	return s
}

// GetMD5Encode 将一个字符串进行MD5加密后返回加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// SimpleCopyFields
//参数传递时，第二个参数使用指针还是实例请自行斟酌，第一个参数必须是指针，涉及的字段必须是对外的
//需要注意的是，该拷贝方法为浅拷贝，换句话说，如果说对象内嵌套有其他的引用类型如Slice,Map等，
//用此方法完成拷贝后，源对象中的引用类型属性内容发生了改变，该对象对应的属性中内容也会改变。
func SimpleCopyFields(dst, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针，.Elem()类似于*ptr的操作返回指针指向的地址反射类型
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	FieldNums := dstType.NumField()

	for i := 0; i < FieldNums; i++ {
		// 属性
		field := dstType.Field(i)
		// 待填充属性值
		fieldValue := srcValue.FieldByName(field.Name) // 无效，说明src没有这个属性 || 属性同名但类型不同
		if !fieldValue.IsValid() || field.Type != fieldValue.Type() {
			continue
		}
		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(fieldValue)
		}
	}

	return nil
}

const (
	BASE    = "E8S2DZX9WYLTN6BQF7CP5IK3MJUAR4HV"
	DECIMAL = 32
	PAD     = "A"
	LEN     = 8
)

// Encode id转code
func Encode(uid uint64) string {
	id := uid
	mod := uint64(0)
	res := ""
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		res += string(BASE[mod])
	}
	resLen := len(res)
	if resLen < LEN {
		res += PAD
		for i := 0; i < LEN-resLen-1; i++ {
			res += string(BASE[(int(uid)+i)%DECIMAL])
		}
	}
	return res
}

func Decode(code string) uint64 {
	res := uint64(0)
	lenCode := len(code)
	baseArr := []byte(BASE)       // 字符串进制转换为byte数组
	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}
	// 查找补位字符的位置
	isPad := strings.Index(code, PAD)
	if isPad != -1 {
		lenCode = isPad
	}
	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == PAD {
			continue
		}
		index, ok := baseRev[code[i]]
		if !ok {
			return 0
		}
		b := uint64(1)
		for j := 0; j < r; j++ {
			b *= DECIMAL
		}
		res += uint64(index) * b
		r++
	}
	return res
}

// PrintPanicStack 产生panic时的调用栈打印
func PrintPanicStack(extras ...interface{}) {
	i := 0
	funcName, file, line, ok := runtime.Caller(i)
	for ok {
		logging.Logger.Debugf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
		i++
		funcName, file, line, ok = runtime.Caller(i)
	}
	for k := range extras {
		logging.Logger.Debugf("EXRAS#%v DATA:%v", k, spew.Sdump(extras[k]))
	}
}

func Serialization(obj interface{}) string {
	data, _ := json.Marshal(obj)
	return string(data)
}

// Intersect 求交集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			logging.Logger.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				logging.Logger.Error("create directory"+v, " error:", err)
				return err
			}
		}
	}
	return err
}

func GetCurrentRow(sheetMap map[string]int64, sheetName string) (bool, int64) {
	i, ok := sheetMap[sheetName]
	if ok {
		//存在即返回value
		return ok, i
	} else {
		//不存在即添加
		sheetMap[sheetName] = 0
		return ok, 0
	}
}

const (
	E    = 0
	D    = 25
	C    = 50
	B    = 75
	A    = 100
	NULL = -1
)

func GetScore(rating string) int {
	switch rating {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	case "D":
		return D
	case "E":
		return E
	default:
		return NULL
	}
}
