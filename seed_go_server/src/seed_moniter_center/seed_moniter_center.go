package main

import (
	"seedetcd"
	"seedlog"
)

type SeedMonServer struct {
	mClient  seedetcd.V3
	mLog     seedlog.SeedLog
	mConfig  SeedMonConfig
}

func (base* SeedMonServer) GetConfig() error {

	return nil
}

func (base* SeedMonServer) Main() {


}






