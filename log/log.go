package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

//定义日志等级类型
type LEVEL byte

const DATE_FORMAT = "2006-01-02"

var fileLogger *FileLogger

const (
	DEBUG LEVEL = iota //0
	INFO               //1
	WARN               //2
	ERROR              //3
)

type FileLogger struct {
	fileDir        string        //日志存储路径
	fileName       string        //日志文件名称
	prefix         string        //日志消息前缀
	logLevel       LEVEL         //日志等级
	logFile        *os.File      //日志文件
	date           *time.Time    //日志当前时间
	lg             *log.Logger   //系统日志对象
	mu             *sync.RWMutex //读写锁，在进行日志分割和日志写入时需要锁住
	logChan        chan string   //日志消息通道，以实现异步写日志
	stopTickerChan chan bool     //停止定时器的通道
}

//初始化系统日志
func Init(fileDir, fileName, prefix, level string) error {
	CloseLogger()

	//初始化日志对象
	f := &FileLogger{
		fileDir:        fileDir,
		fileName:       fileName,
		prefix:         prefix,
		mu:             new(sync.RWMutex),
		logChan:        make(chan string, 5000),
		stopTickerChan: make(chan bool, 1),
	}

	//定义日志等级
	switch strings.ToUpper(level) {
	case "DEBUG":
		f.logLevel = DEBUG
	case "WARN":
		f.logLevel = WARN
	case "ERROR":
		f.logLevel = ERROR
	default:
		f.logLevel = INFO
	}

	t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
	f.date = &t
	f.isExistOrCreateFileDir()

	fullFileName := filepath.Join(f.fileDir, f.fileName+".log")
	file, err := os.OpenFile(fullFileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	f.logFile = file
	f.lg = log.New(f.logFile, prefix, log.LstdFlags|log.Lmicroseconds)

	go f.logWriter()
	go f.fileMonitor()
	fileLogger = f
	return nil

}

// 关闭文件，关闭通道，为了停止一个不断循环的 goroutine
// 由于初始化函数可以会被调用多次，以实现配置的变更，如果不先关闭结束旧的goroutine
// 那同样功能的goroutine 将不止一个在同时运行
func CloseLogger() {
	if fileLogger != nil {
		fileLogger.stopTickerChan <- true
		close(fileLogger.stopTickerChan)
		close(fileLogger.logChan)
		fileLogger.lg = nil
		fileLogger.logFile.Close()
	}
}

//判断日志目录是否存在，不存在则创建
func (this *FileLogger) isExistOrCreateFileDir() {
	_, err := os.Stat(this.fileDir)
	if os.IsNotExist(err) {
		os.Mkdir(this.fileDir, os.ModePerm)
	}
}

// 将日志消息写入文件
func (this *FileLogger) logWriter() {
	defer func() { recover() }()

	for {
		str, ok := <-this.logChan
		if !ok {
			return
		}
		this.mu.RLock()
		this.lg.Output(2, str)
		this.mu.RUnlock()
	}
}

func (this *FileLogger) fileMonitor() {
	defer func() {
		recover()
	}()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if this.isMustSplit() {
				if err := this.split(); err != nil {
					Error("Log split error: %v\n", err.Error())
				}
			}
		}
	}
}

// 判断文件是否需要分割
func (this *FileLogger) isMustSplit() bool {
	t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
	return t.After(*this.date)
}

//日志分割
func (this *FileLogger) split() error {
	this.mu.Lock()
	defer this.mu.Unlock()

	logFile := filepath.Join(this.fileDir, this.fileName)
	//日志备份
	logFileBak := logFile + "-" + this.date.Format(DATE_FORMAT) + ".log"
	if this.logFile != nil {
		this.logFile.Close()
	}

	err := os.Rename(logFile, logFileBak)
	if err != nil {
		return err
	}

	t, _ := time.Parse(DATE_FORMAT, time.Now().Format(DATE_FORMAT))
	this.date = &t
	this.logFile, err = os.OpenFile(logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	this.lg = log.New(this.logFile, this.prefix, log.LstdFlags|log.Lmicroseconds)
	return nil
}

func Info(format string, v ...interface{}) {
	//_, file, line, _ := runtime.Caller(1)
	if fileLogger.logLevel <= INFO {
		//fileLogger.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprintf("[INFO]"+format, v...)
		fileLogger.logChan <- fmt.Sprintf("[INFO]"+format, v...)
		fmt.Println(fmt.Sprintf("[INFO]"+format, v...))
	}
}

func Error(format string, v ...interface{}) {
	//_, file, line, _ := runtime.Caller(1)
	if fileLogger.logLevel <= ERROR {
		//文件行号+日志级别
		//fileLogger.logChan <- fmt.Sprintf("[%v:%v]", filepath.Base(file), line) + fmt.Sprintf("[INFO]"+format, v...)
		fileLogger.logChan <- fmt.Sprintf("[INFO]"+format, v...)
	}
}

func Debug(format string, v ...interface{}) {
	if fileLogger.logLevel <= DEBUG {
		fmt.Printf("[DEBUG]"+format, v...)
	}
}

func Warn(format string, v ...interface{}) {
	if fileLogger.logLevel <= WARN {
		fileLogger.logChan <- fmt.Sprintf("[WARN]"+format, v...)
	}
}

//初始化日志
func InitLog() {
	Init("logs", "kunkka-match", "[KUNKKA] –– ", "info")
}
