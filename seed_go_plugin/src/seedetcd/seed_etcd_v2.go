package seedetcd

import "time"

type V2 struct {
	//[]string{"localhost:2379", "localhost:22379", "localhost:32379"}
	EPath  []string
	DialTimeout time.Duration
	RequestTimeout time.Duration
}