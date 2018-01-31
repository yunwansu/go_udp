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

		udpConn, err1 := net.ListenUDP("udp", udp)
		if err != nil {
			fmt.Printf("Listen UDP Error ~> %d\n", err1)
		}

		data := make([]byte, 1024)

		for {
			length, err3 := udpConn.Read(data)
			if err != nil {
				fmt.Printf("Read Error ~> %d\n", err3)
			}

			fmt.Printf("read data length ~> %d\n", length)
			fmt.Printf("read data ~> %s\n", string(data))
		}

		err2 := udpConn.Close()
		if err != nil {
			fmt.Printf("Close UDP Error ~> %d\n", err2)
		}
}
