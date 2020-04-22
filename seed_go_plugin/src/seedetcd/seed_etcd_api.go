package seedetcd

import (
	"../seedcomdata"
	"../seedlog"
	"encoding/json"
	"time"
)

type UserApp interface {

	GetVersion() string
	Del  (key string)
	Set  (key string, json string) error
	Gets (key string) ([] seedcomdata.SeedEtcdResp, error)
}

const SEED_ETCD_VERSION = "1.0.0.0"

type V3 struct {
	//[]string{"localhost:2379", "localhost:22379", "localhost:32379"}
	EPath  []string
	DialTimeout time.Duration
	RequestTimeout time.Duration
	Log seedlog.SeedLog
}

func (base*V3) GetVersion() string {

	data, _ := json.Marshal(struct {
		VersionType  string
		VersionValue string
	}{
		VersionType:"SEED_ETCD_VERSION",
		VersionValue:SEED_ETCD_VERSION,
	})

	return string(data)
}

