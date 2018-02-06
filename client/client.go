package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type ServerInfo struct {
	Cpu string
	Memory string
	Network string
}

func main() {
	udp, err := net.ResolveUDPAddr("udp", "127.0.0.1:4885")
	if err != nil {
		fmt.Printf("Create UDP Error ~> %d\n", err)
	}

	udpConn, err1 := net.DialUDP("udp", nil, udp)
	if err1 != nil {
		fmt.Printf("Dial UDP Error ~> %d\n", err1)
	}

	serverInfo := ServerInfo{Cpu : "Intel® Xeon® Platinum 8180M Processor",
							 Memory : "PC3-12800",
						 	 Network : "RTL8111D"}

	data, err := json.Marshal(serverInfo)
	// data := []byte("hello UDP!")
	if err != nil {
		fmt.Println("Error : ", err)
	}

	os.Stdout.Write(data)
	udpConn.Write(data)

	err2 := udpConn.Close()
	if err2 != nil {
		fmt.Printf("Close UDP Error ~> %d\n", err2)
	}
}
