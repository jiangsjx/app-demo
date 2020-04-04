package kit

import (
	"log"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog(config *Config) {
	if config.Log.Debug {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: "06-01-02 15:04:05.00",
		})
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		dir := path.Dir(config.Log.Path)
		if _, err := os.Stat(dir); err != nil {
			if !os.IsNotExist(err) {
				log.Fatalf("Stat log directory failed: %s", err.Error())
			}

			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				log.Fatalf("Create log directory failed: %s", err.Error())
			} else {
				log.Printf("Create log directory success")
			}
		}

		jack := &lumberjack.Logger{
			Filename:   config.Log.Path,
			MaxSize:    16, // MB
			MaxBackups: 8,
			MaxAge:     7, // Days
			LocalTime:  true,
			Compress:   true,
		}
		logrus.SetOutput(jack)

		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors:   true,
			FullTimestamp:   true,
			TimestampFormat: "06-01-02 15:04:05.00",
		})
		logrus.SetLevel(logrus.InfoLevel)
	}
}
