package main

import (
	"flag"
	"fmt"
	"github.com/djwackey/dorsvr/rtspclient"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Please input rtsp url.")
		return
	}

	//rtsp_url := os.Args[1]

	rtsp_url := "rtsp://admin:Admin123456@10.88.192.208"

	client := rtspclient.New()

	// to connect rtsp server
	if !client.DialRTSP(rtsp_url) {
		return
	}

	// send the options/describe request
	client.SendRequest()

	select {}
}
