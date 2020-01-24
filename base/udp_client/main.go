package main

import (
	"fmt"
	"net"
)

func main() {
	//connect server
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	})

	if err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}

	//send data
	_, err = conn.Write([]byte("hello server!"))
	if err != nil {
		fmt.Printf("send data failed, err : %v\n", err)
		return
	}

	//receive data from server
	result := make([]byte, 4096)
	n, remoteAddr, err := conn.ReadFromUDP(result)
	if err != nil {
		fmt.Printf("receive data failed, err: %v\n", err)
		return
	}
	fmt.Printf("receive from addr: %v  data: %v\n", remoteAddr, string(result[:n]))
}