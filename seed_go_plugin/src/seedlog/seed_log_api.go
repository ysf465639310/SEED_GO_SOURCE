package seedlog

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

const SEED_LOG_VERSION  = "1.0.0.0"

type SeedLog struct {

	//FileName  string  `json:"FileName"`
	FilePath   string `json:"FilePath"`
	JsonFormat bool `json:"JsonFormat"`
	LogLevel   log.Level `json:"LogLevel"`
	mLog       *log.Logger
}

func (base* SeedLog) GetVersion() string {

	data, _ := json.Marshal(struct {
		VersionType  string
		VersionValue string
	}{
		VersionType:"SEED_LOG_VERSION",
		VersionValue:SEED_LOG_VERSION,
	})

	return string(data)
}

/*

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

*/


