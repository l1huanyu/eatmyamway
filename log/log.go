package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var gFile *os.File

func Init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	gFile, err := os.OpenFile(viper.GetString("log_file"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(gFile)
}

func Close() {
	gFile.Close()
}

func Info(funcName string, msg string, fields map[string]interface{}) {
	logrus.WithField("FuncName", funcName).WithFields(fields).Info(msg)
}

func Warn(funcName string, msg string, fields map[string]interface{}) {
	logrus.WithField("FuncName", funcName).WithFields(fields).Warn(msg)
}

func Error(funcName string, msg string, fields map[string]interface{}) {
	logrus.WithField("FuncName", funcName).WithFields(fields).Error(msg)
}
