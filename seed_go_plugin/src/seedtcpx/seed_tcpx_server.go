package seedtcpx

import (
	"fmt"
	"github.com/fwhezfwhez/tcpx"
	log "github.com/sirupsen/logrus"
	"seedlog"
	"time"
)

// OnConnect is a event callback when a connection is built
func (base* SeedTXServer) OnConnect(c *tcpx.Context) {
	fmt.Println(fmt.Sprintf("connecting from remote host %s network %s", c.ClientIP(), c.Network()))
}
// OnClose is a event callback when a connection is closing
func (base* SeedTXServer) OnClose(c *tcpx.Context) {
	fmt.Println(fmt.Sprintf("connecting from remote host %s network %s has stoped", c.ClientIP(), c.Network()))
}

func (base* SeedTXServer) ServerInit(addr string, protocol string, handles []TxHandle) error {

	if base.log == nil {
		base.log=new(seedlog.SeedLog)
		_ = base.log.Init()
	}

	base.context = tcpx.NewTcpX(tcpx.JsonMarshaller{})

	if base.onConnect == nil {
		base.context.OnConnect = base.OnConnect
		
	} else {
		base.context.OnConnect = base.onConnect
	}

	if base.onClose == nil {
		base.context.OnClose = base.OnClose

	} else {
		base.context.OnClose = base.onClose
	}

	base.context.HeartBeatMode(true, 5*time.Second)

	for _, item := range handles  {
		base.context.AddHandler(item.ID,item.Handle)
	}

	go func() {
		//fmt.Println("tcp srv listen on 7171")
		e := base.context.ListenAndServe(protocol, addr)

		if  e != nil {
			//panic(e)
			base.log.GetLogHandle().WithFields(log.Fields{"err":e}).Info("fail create server...")
		}

	}()

	return nil
}


func (base* SeedTXServer) SetSpecHandle(onClose TxHandler, onConnect TxHandler) {

	base.onConnect = onConnect
	base.onClose = onClose
	return
}




