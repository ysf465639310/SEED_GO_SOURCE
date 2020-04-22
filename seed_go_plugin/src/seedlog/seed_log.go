package seedlog

import (
	rotelog "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func (base* SeedLog) Init() error{
	base.mLog = log.New()

	if base.JsonFormat {
		base.mLog.SetFormatter(&log.JSONFormatter{})
	} else {
		base.mLog.SetFormatter(&log.TextFormatter{})
	}

	if len(base.FilePath) == 0 {
		base.mLog.SetOutput(os.Stdout)
	} else {
		writer, _ := rotelog.New(
			base.FilePath+".%Y%m%d%H%M",
			rotelog.WithLinkName(base.FilePath),
			rotelog.WithMaxAge(time.Duration(7*24)*time.Hour),
			rotelog.WithRotationTime(time.Duration(24)*time.Hour),
			rotelog.WithRotationCount(4),
		)
		log.SetOutput(writer)
	}

	if base.LogLevel < 0 || base.LogLevel > 6 {
		base.mLog.SetLevel(log.DebugLevel)

	} else {
		log.SetLevel(base.LogLevel)
	}

	return nil
}

func (base* SeedLog) GetLogHandle() *log.Logger {

	if base.mLog == nil {
		base.mLog = log.New()
		base.Init()
	}
	return base.mLog
}

