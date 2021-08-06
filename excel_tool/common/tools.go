package common

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"excel_tool/logging"
	"github.com/jordan-wright/email"
	"github.com/oklog/ulid/v2"
	"math"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func SHA256(str string) string {
	m := sha256.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

func ULID() string {
	t := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func RemoveRepByMap(slc []string) []string {
	result := []string{}         //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func GetKeysInt(m map[string]int) []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率较高
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

type (
	//Queue 队列
	Queue struct {
		top    *node
		rear   *node
		length int
	}
	//双向链表节点
	node struct {
		pre   *node
		next  *node
		value interface{}
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

//获取队列长度
func (this *Queue) Len() int {
	return this.length
}

//返回true队列不为空
func (this *Queue) Any() bool {
	return this.length > 0
}

//返回队列顶端元素
func (this *Queue) Peek() interface{} {
	if this.top == nil {
		return nil
	}
	return this.top.value
}

//入队操作
func (this *Queue) Push(v interface{}) {
	n := &node{nil, nil, v}
	if this.length == 0 {
		this.top = n
		this.rear = this.top
	} else {
		n.pre = this.rear
		this.rear.next = n
		this.rear = n
	}
	this.length++
}

//出队操作
func (this *Queue) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	if this.top.next == nil {
		this.top = nil
	} else {
		this.top = this.top.next
		this.top.pre.next = nil
		this.top.pre = nil
	}
	this.length--
	return n.value
}

func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

func InterfaceToString(json interface{}) string {
	var s string
	switch json.(type) {
	case nil:
		s = ""
	case string:
		s = json.(string)
	case float64:
		f := json.(float64)
		s = strconv.FormatFloat(f, 'f', 0, 64)
	default:
		panic("类型不能识别")
	}
	return s
}

func InterfaceToInt(json interface{}) int {
	var i int
	switch json.(type) {
	case nil:
		i = 0
	case string:
		i, _ = strconv.Atoi(json.(string))
	case float64:
		f := json.(float64)
		i = int(f)
	default:
		panic("类型不能识别")
	}
	return i
}

func SliceFind(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func CheckUidList(collection []string) bool {
	if len(collection) == 0 {
		return false
	}
	for _, uid := range collection {
		if len(strings.TrimSpace(uid)) != 20 {
			return false
		}
	}
	return true
}

var gracefulWaitGroup sync.WaitGroup

func GracefulWorkerAdd(delta int) {
	gracefulWaitGroup.Add(delta)
}

func GracefulWorkerDone() {
	gracefulWaitGroup.Done()
}

func GracefulWorkerWait() {
	gracefulWaitGroup.Wait()
}

func WaitWorker() {
	logging.Logger.Debug("waiting worker")
	gracefulWaitGroup.Wait()
	logging.Logger.Debug("all worker is done.")
}

// MailAlarm 邮件报警
func MailAlarm(message string) {
	//将验证码通过邮件发送给用户
	mail := email.NewEmail()
	//设置发送方的邮箱
	mail.From = Sender
	// 设置接收方的邮箱
	mail.To = []string{Sender}
	//设置主题
	mail.Subject = EmailSubject
	//设置文件发送的内容
	mail.Text = []byte("服务异常：" + message)
	//设置服务器相关的配置
	err := mail.Send(EmailServerAddr, smtp.PlainAuth("", Sender, EmailAuthorizationCode, EmailHost))
	if err != nil {
		logging.Logger.Error(err)
		return
	}
}
