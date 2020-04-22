package main

import (
	log "github.com/sirupsen/logrus"
	_ "go.etcd.io/etcd/mvcc/backend"
	"seedcomdata"
	"seedetcd"
	"seedlog"
	"seedtcpx"
	"time"
)

type SeedChatServer struct {
	mClient  seedetcd.V3
	mLog     seedlog.SeedLog
	mConfig  SeedChatConfig
	mServer  seedtcpx.SeedTXServer
	mStop    chan string
	mIsStop  bool
	handles  []seedtcpx.TxHandle
}


func (base* SeedChatServer) getCfg() error {
	return nil
}

func (base* SeedChatServer) logInit() error {
	if len(base.mConfig.Path) == 0 && base.mConfig.IsRelease == false{
		base.mConfig.Path="./log/seed_center"
	}
	base.mLog.FilePath=base.mConfig.Path
	return  base.mLog.Init()
}

func (base* SeedChatServer) stop()  {
	//recv stop signal
	close(base.mStop)
	return
}

func (base* SeedChatServer) wait()  {

	select {
	case _ = <- base.mStop:
		base.mLog.GetLogHandle().WithFields(log.Fields{}).Info("server will stop wait 5s...")
		base.mIsStop = true
	}
	time.Sleep(5)
	return
}

func (base *SeedChatServer) TxServerInit() {
	var path string = ":" + string(base.mConfig.Port)
	base.mServer.ServerInit(path, "tcp", base.handles)
}

func (base *SeedChatServer) busInit ()  {

	base.handles = append(base.handles,
		seedtcpx.TxHandle{ID:int32(seedcomdata.REQ_ONLINE), Handle:base.Online})
	base.handles = append(base.handles,
		seedtcpx.TxHandle{ID:int32(seedcomdata.REQ_OFFLINE),Handle:base.Offline})
	base.handles = append(base.handles,
		seedtcpx.TxHandle{ID:int32(seedcomdata.REQ_SENDMS), Handle:base.ReqSendM})
	base.handles = append(base.handles,
		seedtcpx.TxHandle{ID:int32(seedcomdata.REQ_CRHOME), Handle:base.CRoom})
	return

}

func (base *SeedChatServer) CreateRoom(login seedcomdata.Login)  {

	return
}

func (base* SeedChatServer) Main() {
	base.mStop=make(chan string)
	err := base.getCfg()
	if err != nil {
		base.stop()
		goto STOP
	}
	err = base.logInit()
	if err != nil {
		base.stop()
		goto STOP
	}

	base.busInit()

	base.TxServerInit()
STOP:
	base.wait()
	return
}