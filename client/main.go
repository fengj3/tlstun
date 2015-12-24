package main

import (
	"flag"
	"fmt"
)

var listenIp string
var listenPort int

var serverIp string
var serverPort int

func main() {
	Log("daemon", "info", "starting proxy")
	flag.Parse()
	wsServer := fmt.Sprintf("%s:%d", serverIp, serverPort)
	forward(wsServer)
}

func init() {
	flag.IntVar(&listenPort, "port", 1080, "port to listen on")
	flag.StringVar(&listenIp, "ip", "127.0.0.1", "ip to bind to")
	flag.IntVar(&serverPort, "sport", 12345, "port to listen on")
	flag.StringVar(&serverIp, "sip", "62.148.169.249", "ip to bind to")
}
