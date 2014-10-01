package main

import (
	"net"
)

//var block chan int = make(chan int)

func myUDPServer() {
	udpAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:80")
	if err != nil {
		panic(err)
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
//	read := []byte{}
	read := make([]byte, 1500)
	for {
		_, _, err = udpConn.ReadFrom(read)
		if err != nil {
			panic(err)
		}
		println("read content:", string(read))
	}
}

func main() {
	println("^^^^running")
	myUDPServer()
}
