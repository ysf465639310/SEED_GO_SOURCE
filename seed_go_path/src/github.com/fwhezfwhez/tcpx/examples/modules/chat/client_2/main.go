package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fwhezfwhez/tcpx"
	"net"
)

var panel = make(chan string, 100)
var conn net.Conn
var user = "Li San"
var toUser = "Zhao Qiang"

func init() {
	var e error
	// connect
	conn, e = net.Dial("tcp", "localhost:8103")
	if e != nil {
		panic(e)
	}
	// receive from server
	go Recv(conn)

	// spying output panel
	go func() {
		for {
			select {
			case msg := <-panel:
				fmt.Print(msg)
			}
		}
	}()

	// online
	online(user, conn)
}
func main() {
	f := bufio.NewReader(os.Stdin)
	for {
		input, _ := f.ReadString('\n')
		if len(input) == 1 {
			continue
		}
		send(input, toUser, conn)
	}
	select {}
}

func online(username string, conn net.Conn) {
	onlineBuf, e := tcpx.PackWithMarshaller(tcpx.Message{
		MessageID: 1,
		Header:    nil,
		Body:      struct{ Username string `json:"username"` }{username},
	}, tcpx.JsonMarshaller{})
	if e != nil {
		panic(e)
	}
	conn.Write(onlineBuf)
	panel <- "online success"
}

func Recv(conn net.Conn) {
	var buf = make([]byte, 500)
	var e error
	for {
		buf, e = tcpx.FirstBlockOf(conn)
		if e != nil {
			fmt.Println(e.Error())
			os.Exit(0)
			break
		}
		bf, _ := tcpx.BodyBytesOf(buf)
		messageID, e := tcpx.MessageIDOf(buf)

		switch messageID {
		case 400, 200, 500:
			type ResponseTo struct {
				Message string `json:"message"`
			}
			var rs ResponseTo
			e = json.Unmarshal(bf, &rs)
			if e != nil {
				panic(e)
			}
			panel <- fmt.Sprintf("message: %s", rs.Message)
		case 6:
			type ResponseTo struct {
				Message  string `json:"message"`
				FromUser string `json:"from_user"`
			}
			var rs ResponseTo
			e = json.Unmarshal(bf, &rs)
			if e != nil {
				panic(e)
			}
			panel <- fmt.Sprintf("%s: %s", rs.FromUser, rs.Message)
		}

	}
}

func send(msg string, toUser string, conn net.Conn) {
	buf, e := tcpx.PackWithMarshaller(tcpx.Message{
		MessageID: 5,
		Header:    nil,
		Body: struct {
			Message string `json:"message"`
			ToUser  string `json:"to_user"`
		}{Message: msg, ToUser: toUser},
	}, tcpx.JsonMarshaller{})
	if e != nil {
		panic(e)
	}
	conn.Write(buf)
}
