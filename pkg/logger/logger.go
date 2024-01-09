package logger

import (
	"example.com/v2/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"
)

func init() {
	//设置日志格式的json格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func Write(msg string, filename string) {
	setOutputPutfile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutputPutfile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)
}
func Info(fields logrus.Fields, args ...interface{}) {
	setOutputPutfile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Debug(args)
}
func Warn(fields logrus.Fields, args ...interface{}) {
	setOutputPutfile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Debug(args)
}
func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutputPutfile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Debug(args)
}
func Error(fields logrus.Fields, args ...interface{}) {
	setOutputPutfile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Debug(args)
}
func Panic(fields logrus.Fields, args ...interface{}) {
	setOutputPutfile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Debug(args)
}
func Trace(fields logrus.Fields, args ...interface{}) {
	setOutputPutfile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Debug(args)
}
func setOutputPutfile(level logrus.Level, logName string) {
	Logdir := config.Logdir
	timeStr := time.Now().Format("2006-01-02")
	logpath := Logdir + timeStr
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		err = os.MkdirAll(logpath, 0777)
		if err != nil {
			panic(fmt.Errorf("creat log dir '%s'error: %s", "./runtime/log", err))
		}
	}

	fileName := path.Join(logpath, logName+"_"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file err", err)
	}
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

func LoggerToFile() gin.LoggerConfig {
	Logdir := config.Logdir
	timeStr := time.Now().Format("2006-01-02")
	logpath := Logdir + timeStr
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		err = os.MkdirAll(logpath, 0777)
		if err != nil {
			panic(fmt.Errorf("creat log dir '%s'error: %s", "./runtime/log", err))
		}
	}
	fileName := path.Join(logpath, "success_"+timeStr+".log")

	os.Stderr, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \" %s %s %s %d %s \"  %s  %s\n",
				params.TimeStamp.Format("2006-01-02 15:04:45"),
				params.ClientIP,
				params.Method,
				params.Request.URL,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)

		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}
	return conf
}

func Recover(c *gin.Context) {
	Logdir := config.Logdir
	timeStr := time.Now().Format("2006-01-02")
	logpath := Logdir + timeStr
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat(logpath); os.IsNotExist(errDir) {
				errDir = os.MkdirAll(logpath, 0777)
				if err != nil {
					panic(fmt.Errorf("creat log dir '%s'error: %s", "./runtime/log", err))
				}
			}

			timeStr := time.Now().Format("2006-01-02")
			fileName := path.Join(logpath, "success_"+timeStr+".log")
			f, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if errFile != nil {
				fmt.Println(errFile)
			}
			timeFileStr := time.Now().Format("2006-01-02 15:34:45")
			f.WriteString("panic error time panic:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stackrace from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})
			//终止后续接口调用，不加的话recover到异常后还继续调用接口力后续代码
			c.Abort()
		}
	}()
	c.Next()
}
