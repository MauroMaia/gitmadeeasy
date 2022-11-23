package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

const FILE_PATH = "/tmp/log.txt"

var Logger = log.StandardLogger()

func init() {
	SetLoggerOptions(true, true)
	Logger.Infoln("Logger init")
}

func SetLoggerOptions(useColor bool, verbose bool) {
	Logger.SetFormatter(&log.TextFormatter{
		DisableColors:          !useColor,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	_ = os.Remove(FILE_PATH)
	file, _ := os.Create(FILE_PATH)

	Logger.SetOutput(file)

	if verbose == true {
		Logger.SetLevel(log.TraceLevel)
	}
}
