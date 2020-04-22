package seedtcpx

import (
	"fmt"
	"github.com/fwhezfwhez/tcpx"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"seedlog"
	"time"
)

func (base* SeedTXClient) AsyConnect(addr string, protocol string, handles []TcHandle) error {

	if base.log == nil {
		base.log=new(seedlog.SeedLog)
		_ = base.log.Init()
	}
	//base.context = tcpx.NewTcpX(tcpx.JsonMarshaller{})
	var er error
	base.con , er= net.Dial(protocol, addr)
	if er != nil {
		return er
	}
	base.handles = handles
	go base.recv()
	go base.heart()
	return nil
}

func (base* SeedTXClient) heart()  {
	var e error
	var heartBeat []byte
	heartBeat, e = tcpx.PackWithMarshaller(tcpx.Message{
		MessageID: tcpx.DEFAULT_HEARTBEAT_MESSAGEID,
		Header:    nil,
		Body:      nil,
	}, nil)

	for {
		_, e = base.con.Write(heartBeat)
		if e != nil {
			//fmt.Println(e.Error())
			break
		}
		time.Sleep(10 * time.Second)
	}
}

func (base* SeedTXClient) recv() {
	var buf = make([]byte, 500)
	var e error
	for  {
		buf, e = tcpx.FirstBlockOf(base.con)
		if e != nil {
			fmt.Println(e.Error())
			os.Exit(0)
		}
		bf, _ := tcpx.BodyBytesOf(buf)
		messageID, _:= tcpx.MessageIDOf(buf)
		base.log.GetLogHandle().WithFields(log.Fields{"messageId":messageID}).Info("recv message ...")
		for _,item := range base.handles{
			if item.ID == messageID {
				item.Handle(base.con,item.ID,string(bf))
			}
		}
	}
}

