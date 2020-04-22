// Package server executable file
package main

import (
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"github.com/fwhezfwhez/tcpx"
	//"tcpx"
)

func main() {
	srv := tcpx.NewTcpX(tcpx.JsonMarshaller{})

	// If mode is DEBUG, error in framework will log with error spot and time in detail
	// tcpx.SetLogMode(tcpx.DEBUG)

	// udp is no-state protocol, no need to have these
	srv.OnClose = nil
	srv.OnConnect = nil

	// Mux routine and OnMessage callback can't meet .
	// When OnMessage is not nil, routes will lose effect.
	// When srv.OnMessage has set, srv.AddHandler() makes no sense, it means user wants to handle raw message stream by self.
	// Besides, if OnMessage is not nil, middlewares of global type(by srv.UseGlobal) and anchor type(by srv.Use, srv.UnUse)
	// will all be executed regardless of an anchor type middleware being unUsed or not.
	// srv.OnMessage = OnMessage

	srv.UseGlobal(MiddlewareGlobal)
	srv.Use("middleware1", Middleware1, "middleware2", Middleware2)
	srv.AddHandler(1, SayHello)

	srv.UnUse("middleware2")
	srv.AddHandler(3, SayGoodBye)

	srv.AddHandler(5, Middleware3, SayName)
	fmt.Println("udp srv listen on 7172")
	if e := srv.ListenAndServe("udp", ":7172"); e != nil {
		panic(e)
	}
}

func OnClose(c *tcpx.Context) {
	fmt.Println(fmt.Sprintf("connecting from remote host %s network %s has stoped", c.ClientIP(), c.Network()))
}

var packx = tcpx.NewPackx(tcpx.JsonMarshaller{})

func OnMessage(c *tcpx.Context) {
	type ServiceA struct {
		Username string `json:"username"`
	}
	type ServiceB struct {
		ServiceName string `json:"service_name" toml:"service_name" yaml:"service_name"`
	}

	messageID, e := packx.MessageIDOf(c.Stream)
	if e != nil {
		fmt.Println(errorx.Wrap(e).Error())
		return
	}

	switch messageID {
	case 7:
		var serviceA ServiceA
		// block, e := packx.Unpack(c.Stream, &serviceA)
		block, e := c.Bind(&serviceA)
		fmt.Println(block, e)
		c.Reply(8, "success")
	case 9:
		var serviceB ServiceB
		//block, e := packx.Unpack(c.Stream, &serviceB)
		block, e := c.Bind(&serviceB)
		fmt.Println(block, e)
		c.JSON(10, "success")
	}

}
func SayHello(c *tcpx.Context) {
	var messageFromClient string
	var messageInfo tcpx.Message
	messageInfo, e := c.Bind(&messageFromClient)
	if e != nil {
		panic(e)
	}
	fmt.Println("receive messageID:", messageInfo.MessageID)
	fmt.Println("receive header:", messageInfo.Header)
	fmt.Println("receive body:", messageInfo.Body)

	var responseMessageID int32 = 2
	e = c.Reply(responseMessageID, "hello")
	fmt.Println("reply:", "hello")
	if e != nil {
		fmt.Println(e.Error())
	}
}

func SayGoodBye(c *tcpx.Context) {
	var messageFromClient string
	var messageInfo tcpx.Message
	messageInfo, e := c.Bind(&messageFromClient)
	if e != nil {
		panic(e)
	}
	fmt.Println("receive messageID:", messageInfo.MessageID)
	fmt.Println("receive header:", messageInfo.Header)
	fmt.Println("receive body:", messageInfo.Body)

	var responseMessageID int32 = 4
	e = c.Reply(responseMessageID, "bye")
	fmt.Println("reply:", "bye")
	if e != nil {
		fmt.Println(e.Error())
	}
}

func SayName(c *tcpx.Context) {
	var messageFromClient string
	var messageInfo tcpx.Message
	messageInfo, e := c.Bind(&messageFromClient)
	if e != nil {
		panic(e)
	}
	fmt.Println("receive messageID:", messageInfo.MessageID)
	fmt.Println("receive header:", messageInfo.Header)
	fmt.Println("receive body:", messageInfo.Body)

	var responseMessageID int32 = 6
	e = c.Reply(responseMessageID, "my name is tcpx")
	fmt.Println("reply:", "my name is tcpx")
	if e != nil {
		fmt.Println(e.Error())
	}
}

func Middleware1(c *tcpx.Context) {
	fmt.Println("I am middleware 1 exampled by 'srv.Use(\"middleware1\", Middleware1)'")
}

func Middleware2(c *tcpx.Context) {
	fmt.Println("I am middleware 2 exampled by 'srv.Use(\"middleware2\", Middleware2),srv.UnUse(\"middleware2\")'")
}

func Middleware3(c *tcpx.Context) {
	fmt.Println("I am middleware 3 exampled by 'srv.AddHandler(5, Middleware3, SayName)'")
}

func MiddlewareGlobal(c *tcpx.Context) {
	fmt.Println("I am global middleware exampled by 'srv.UseGlobal(MiddlewareGlobal)'")
}
