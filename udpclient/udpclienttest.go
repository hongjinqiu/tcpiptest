package main

import (
	"net"
)

//var block chan int = make(chan int)

func myUDPClient() {
	if true {
		var laddr *net.UDPAddr = nil
		raddr, err := net.ResolveUDPAddr("udp", "localhost:80")
		if err != nil {
			panic(err)
		}
		udpConn, err := net.DialUDP("udp", laddr, raddr)
//		_, err = udpConn.WriteTo([]byte("test, from cient"), raddr)
		_, err = udpConn.Write([]byte("test, from cient"))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	myUDPClient()
}
