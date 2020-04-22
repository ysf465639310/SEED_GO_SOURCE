package main

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/tcpx"
	"seedcomdata"
)

const ENV_HOST  = "ENV_HOST"

type SeedChatConfig struct {

	Uid  string    `json:"uid"`
	Name string    `json:"name"`
	Type string    `json:"type"`
	Host string    `json:"host"`
	Port  int	   `json:"port"`
	VirHost string `json:"vir_host"`
	Path string   `json:"path"`
	IsRelease bool  `json:"is_release"`
}

type ROOM struct {
	ID string
}

func (base* SeedChatConfig) CreateJson() (string, error){

	data, err := json.Marshal(base)

	return string(data), err
}


func (base* SeedChatConfig) ParseJson (data string) error {
	return  json.Unmarshal([]byte(data), base)
}



func (base *SeedChatServer) Online(c * tcpx.Context)  {
	var login seedcomdata.Login
	if _, e := c.Bind(&login); e != nil {
		//fmt.Println(errorx.Wrap(e).Error())
		return
	}
	c.Online(login.Username)
}

func (base *SeedChatServer) Offline(c * tcpx.Context)  {
	//fmt.Println("offline success")
	c.Offline()
}

func (base* SeedChatServer) ReqSendM(c * tcpx.Context) {
	var req seedcomdata.ReqSend
	var err  int32 = 0
	var message = ""
	if _, e := c.Bind(&req); e != nil {
		err = 1
		goto STOP
	}

	{
		anotherCtx := c.GetPoolRef().GetClientPool(req.ToUser)
		if anotherCtx.IsOnline() {
			_ = c.SendToUsername(req.ToUser, int32(seedcomdata.REQ_SENDMS), seedcomdata.ReqSend{
				Message:  req.Message,
				FromUser: req.ToUser,
				ToUser:   req.ToUser,
			})
			message = fmt.Sprintf("'%s' succeed ... ", req.ToUser)
			err = 0
			goto STOP

		} else {

			message = fmt.Sprintf("'%s' is offline", req.ToUser)
			err = 1
		}

	}
STOP:

	c.JSON(int32(seedcomdata.REP_SENDMS), seedcomdata.RepSend{
		Err:err,
		Message: message,
		FromUser: req.ToUser,
		ToUser:req.ToUser,
	})

	return
}

func (base* SeedChatServer) CRoom(c * tcpx.Context) {
	var req seedcomdata.Login
	var err int32 = 0
	var message = ""
	if _, e := c.Bind(&req); e != nil {
		err = 1
		goto STOP
	}
	base.CreateRoom(req)
STOP:
	c.JSON(int32(seedcomdata.REP_SENDMS), seedcomdata.Rep{
		Err:err,
		Message: message,
	})
}
