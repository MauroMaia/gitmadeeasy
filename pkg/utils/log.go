package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strconv"
)

const FILE_PATH = "/tmp/log.txt"

var Logger = log.StandardLogger()

func init() {
	SetLoggerOptions(true, true)
	Logger.Infoln("Logger init")
}

func SetLoggerOptions(useColor bool, verbose bool) {

	Logger.ReportCaller = true

	Logger.SetFormatter(&log.TextFormatter{
		DisableColors:          !useColor,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		CallerPrettyfier:       CallerPrettier,
	})

	Logger.SetFormatter(&log.JSONFormatter{
		PrettyPrint:      false,
		DataKey:          "data",
		CallerPrettyfier: CallerPrettier,
	})

	_ = os.Remove(FILE_PATH)
	file, _ := os.Create(FILE_PATH)

	Logger.SetOutput(file)

	if verbose == true {
		Logger.SetLevel(log.TraceLevel)
	}
}

func CallerPrettier(frame *runtime.Frame) (function string, file string) {
	fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
	return function, fileName
}
