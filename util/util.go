package util

import (
	"flag"
	"os"

	"MANAGEMENT-2.0/model"
	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {
	Logger = *logrus.New()
	//Logger.Out = os.Stdout

	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	// Logger.SetFormatter(&logrus.TextFormatter{
	// 	FullTimestamp: true,
	// })
	Logger.SetOutput(os.Stdout)
}
func SetLogger() {
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level(debug , error , info , warnning )")
	flag.Parse()
	switch *logLevel {
	case model.LogLevelDebug:
		Logger.Level = logrus.DebugLevel
	case model.LogLevelError:
		Logger.Level = logrus.ErrorLevel
	case model.LogLevelInfo:
		Logger.Level = logrus.InfoLevel
	default:
		Logger.Level = logrus.WarnLevel

	}
}
func Log(logLevel, packageLevel, functionName string, massage, parametter interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if parametter != nil {
			Logger.Debugf("packegelevel : %s , functionname : %s ,massage : %v ,parametter :%v\n", packageLevel, functionName, massage, parametter)
		} else {
			Logger.Debugf("packegelevel : %s , functionname : %s ,massage : %v\n", packageLevel, functionName, massage)

		}
	case model.LogLevelInfo:
		if parametter != nil {
			Logger.Infof("packegelevel : %s , functionname : %s ,massage : %v ,parametter :%v\n", packageLevel, functionName, massage, parametter)
		} else {
			Logger.Infof("packegelevel : %s , functionname : %s ,massage : %v\n", packageLevel, functionName, massage)

		}
	case model.LogLevelError:
		if parametter != nil {
			Logger.Errorf("packegelevel : %s , functionname : %s ,massage : %v ,parametter :%v\n", packageLevel, functionName, massage, parametter)
		} else {
			Logger.Errorf("packegelevel : %s , functionname : %s ,massage : %v\n", packageLevel, functionName, massage)

		}
	default:
		if parametter != nil {
			Logger.Infof("packegelevel : %s , functionname : %s ,massage : %v ,parametter :%v\n", packageLevel, functionName, massage, parametter)
		} else {
			Logger.Infof("packegelevel : %s , functionname : %s ,massage : %v\n", packageLevel, functionName, massage)

		}
	}
}
