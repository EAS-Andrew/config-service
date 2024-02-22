package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		Log.Warn("Invalid LOG_LEVEL, defaulting to 'info'")
		Log.SetLevel(logrus.InfoLevel)
	} else {
		Log.Infof("Setting log level to '%s'", level)
		Log.SetLevel(level)
	}
}
