package seedtcpx

import (
	"../seedlog"
	"encoding/json"
	"github.com/fwhezfwhez/tcpx"
	"net"
)

const SEED_TX_VERSION = "1.0.0.0"

type TcHandler func (conn net.Conn, ID int32, Message string)

type TcHandle struct {
	ID  int32
	Handle  TcHandler
}

type SeedTXClient struct {
	context *tcpx.TcpX
	log     *seedlog.SeedLog
	con     net.Conn
	handles []TcHandle
}

type TxHandler func (context * tcpx.Context)


type TxHandle struct {
	ID  int32
	Handle  TxHandler
}



func (base* SeedTXClient) GetVersion() string {

	data, _ := json.Marshal(struct {
		VersionType  string
		VersionValue string
	}{
		VersionType:"SEED_TX_VERSION",
		VersionValue:SEED_TX_VERSION,
	})

	return string(data)
}

type SeedTXServer struct {

	context *tcpx.TcpX
	log    *seedlog.SeedLog
	onConnect func (context * tcpx.Context)
	onClose   func(context * tcpx.Context)
}

func (base* SeedTXServer) GetVersion() string {

	data, _ := json.Marshal(struct {
		VersionType  string
		VersionValue string
	}{
		VersionType:"SEED_TX_VERSION",
		VersionValue:SEED_TX_VERSION,
	})

	return string(data)
}

func TXLogInit (level int)  {

	tcpx.SetLogMode(level)

}
