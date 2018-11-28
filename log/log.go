package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var outputDir = "./logs/"

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL

	KB = 1024
	MB = 1024 * KB
)

var (
	Log *Logger
	levers = map[int]string{
		DEBUG : "[DEBUG]:",
		INFO  : "[INFO] :",
		WARN  : "[WARN] :",
		ERROR : "[ERROR]:",
		FATAL : "[FATAL]:",
	}
)

type Logger struct{
	lever   int
	maxSize int64
	file *os.File
	log *log.Logger
	lock sync.Mutex
}

//init create a logger instance
func InitLog(lever int , maxSize int64) {
	f, err := newfile()
	if err != nil || f == nil {
		fmt.Printf("Error init log ")
		os.Exit(-1)
	}
	Log = &Logger{}
	Log.file = f
	Log.lever = lever
	if maxSize == 0 {
		Log.maxSize = 10
	}else {
		Log.maxSize = maxSize
	}
	Log.log = log.New(io.MultiWriter(os.Stdout, Log.file), "", log.Ldate|log.Lmicroseconds)
	go func() {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				overSized := isOverSized(Log)
				if overSized {
					return
				}
			}
		}
	}()
}

func isOverSized(logger *Logger) bool{
	fileInfo , err := logger.file.Stat()
	if err != nil {
		fmt.Println("invalid log stat :", err.Error())
		os.Exit(-1)
	}
	if fileInfo.Size() > logger.maxSize * MB {
		copyLock := &Log.lock
		copyLock.Lock()
		logger.file.Close()
		InitLog(0, logger.maxSize)
		copyLock.Unlock()
		return true
	}
	return false
}

func newfile() (*os.File,error){
	_ , err := os.Stat(outputDir)
	if err != nil {
		err := os.MkdirAll(outputDir,0766)
		if err != nil {
			return nil , err
		}
	}
	name := time.Now().Format("2006-01-02_15.04.05")
	f , err := os.OpenFile(outputDir + name , os.O_RDWR | os.O_CREATE , 0666)
	if err != nil {
		return nil , err
	}
	return f , nil
}

func Debugf(format string, a ...interface{}){
	Log.printMsgf(DEBUG,format ,a...)
}

func Infof(format string , a ...interface{}){
	Log.printMsgf(INFO, format ,a...)
}

func Warnf(format string , a ...interface{}){
	Log.printMsgf(WARN,format,a...)
}

func Errorf(format string ,a ...interface{}){
	Log.printMsgf(ERROR,format ,a...)
}

func Fatalf(format string ,a ...interface{}){
	Log.printMsgf(FATAL,format ,a...)
}


func Debug( a ...interface{}){
	Log.printMsg(DEBUG,a...)
}

func Info(a ...interface{}){
	Log.printMsg(INFO,a...)
}

func Warn(a ...interface{}){
	Log.printMsg(WARN,a...)
}

func Error(a ...interface{}){
	Log.printMsg(ERROR,a...)
}

func Fatal(a ...interface{}){
	Log.printMsg(FATAL,a...)
}

func (logger *Logger) printMsg(lever int , a ...interface{}){
	if logger.lever <= lever {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		logger.log.Output(2,levers[lever]+fmt.Sprint(a...))
	}
}

func (logger *Logger) printMsgf(lever int ,format string, a ...interface{}){
	if logger.lever <= lever {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		logger.log.Output(2,levers[lever]+fmt.Sprintf(format,a...))
		if lever == FATAL {
			os.Exit(-1)
		}
	}
}