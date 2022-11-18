package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

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

	file, _ := os.Create("/tmp/log.txt")

	Logger.SetOutput(file)

	if verbose == true {
		Logger.SetLevel(log.DebugLevel)
	}
}
