// Package client executable file
package main

import (
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"github.com/fwhezfwhez/tcpx"
	"net"
	//"tcpx"
)

var packx = tcpx.NewPackx(tcpx.JsonMarshaller{})
// var packxProto = tcpx.NewPackx(tcpx.ProtobufMarshaller{})

func main() {
	conn, err := net.Dial("tcp", "localhost:7171")

	if err != nil {
		panic(err)
	}
	go func() {
		for {
			buf, e := tcpx.FirstBlockOf(conn)
			if e != nil {
				//if e == io.EOF {
				//	break
				//}
				panic(errorx.Wrap(e))
			}
			fmt.Println(buf)
			var receivedString string
			//fmt.Println(buf)
			message, e := packx.Unpack(buf, &receivedString)
			if e != nil {
				panic(errorx.Wrap(e))
			}
			fmt.Println("收到服务端消息块:", smartPrint(message))
			fmt.Println("服务端消息:", receivedString)
			//var resp pb.SayHelloReponse
			//message, e := tcpx.UnpackWithMarshallerName(buf, &resp, "protobuf")
			//if e != nil {
			//	panic(errorx.Wrap(e))
			//}
			//fmt.Println("收到服务端消息块:", smartPrint(message))
			//fmt.Println("服务端消息:", resp)
		}
	}()

	var buf []byte
	var e error
	//buf, e = tcpx.PackWithMarshallerName(tcpx.Message{
	//	MessageID: 11,
	//	Body: &pb.SayHelloRequest{
	//		Username: "ft",
	//	},
	//}, "protobuf")
	//if e != nil {
	//	panic(e)
	//}
	//go func() {
	//	heartbeat, e := tcpx.PackWithMarshallerName(tcpx.Message{
	//		MessageID: 20,
	//		Body:      nil,
	//	}, "protobuf")
	//	if e != nil {
	//		panic(e)
	//	}
	//	for {
	//		time.Sleep(5 * time.Second)
	//		conn.Write(heartbeat)
	//	}
	//}()
	//_ = buf
	//conn.Write(buf)

	buf, e = packx.Pack(5, "hello,I am client xiao ming", map[string]interface{}{
		"api": "/tcpx/client1/",
	})
	if e != nil {
		panic(e)
	}

	fmt.Println(buf)

	for i:=0;i<1000;i++ {
		conn.Write(buf)
	}

	//buf, e = packx.Pack(7, struct {
	//	Username string `json:"username"`
	//}{"FT"}, map[string]interface{}{
	//	"api": "/tcpx/client1/",
	//})
	//if e != nil {
	//	panic(e)
	//}
	//
	// _=buf
	//conn.Write(buf)

	//buf, e := tcpx.PackWithMarshallerName(tcpx.Message{
	//	MessageID: 9,
	//	Header: map[string]interface{}{
	//		"api": "/tcpx/client1/",
	//	},
	//	Body: struct {
	//		ServiceName string `json:"service_name"`
	//	}{"FT"},
	//}, "json")
	//
	//if e != nil {
	//	panic(e)
	//}
	//conn.Write(buf)

	select {}
}

func smartPrint(src interface{}) string {
	buf, _ := json.MarshalIndent(src, "  ", "  ")
	return string(buf)
}
