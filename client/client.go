package main

import (
	"fmt"
	"net"
)

func main() {
	udp, err := net.ResolveUDPAddr("udp", "127.0.0.1:4885")
	if err != nil {
		fmt.Printf("Create UDP Error ~> %d\n", err)
	}

	udpConn, err1 := net.DialUDP("udp", nil, udp)
	if err1 != nil {
		fmt.Printf("Dial UDP Error ~> %d\n", err1)
	}

	data := []byte("hello UDP!")


	udpConn.Write(data)

	err2 := udpConn.Close()
	if err2 != nil {
		fmt.Printf("Close UDP Error ~> %d\n", err2)
	}
}
