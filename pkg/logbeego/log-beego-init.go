package logbeego

import (
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"
)

// logDir:日志文件夹
func LogBeegoInit(logName string, maxDays int) error {
	if logName == "" {
		err := errors.New("beego日志文件名为空")
		panic(err)
	}
	if !strings.Contains(logName, ".log") {
		logName += ".log"
	}
	if !strings.Contains(logName, "logs/") {
		logName = "logs/" + logName
	}
	if maxDays <= 0 {
		maxDays = 10
	}
	logBuf := fmt.Sprintf(`{"filename":"%s","maxdays":%d}`, logName, maxDays)
	logs.SetLogger(logs.AdapterFile, logBuf)
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(3)
	return nil
}
