package Server

import (
	"fmt"
	"net"
	"strings"
)

func Server(hostname string) {
	server, error := net.ResolveUDPAddr("udp4", hostname)
	if error != nil {
		fmt.Println("Server: (WARNING/FATAL)  -> Server could not resolve or recognise address and hostname -> ", error)
	} else {
		fmt.Println("Server: (Success) -> Server resolved hostname")
	}

	connect, error := net.ListenUDP("udp4", server)
	if error != nil {
		fmt.Println("Server: (WARNING/FATAL) -> Server could not listen on the parsed address -> ", error)
	} else {
		fmt.Println("Server: (Success) -> Server listening on the parsed address")
	}
	defer connect.Close()
	buffer := make([]byte, 1024)
	for {
		network, address, error := connect.ReadFromUDP(buffer)
		fmt.Print("Server got payload -> ", string(buffer[0:network-1]))
		if strings.TrimSpace(string(buffer[0:network])) == "ServerStop" {
			fmt.Println("Shutting down server...")
			return
		}
		msg := "ServerResponse: Got message from client -> " + string(buffer[0:network])
		_, error = connect.WriteToUDP([]byte(msg), address)
		if error != nil {
			fmt.Println("Server: (Failure) -> Server got a weird error -> ", error)
			return
		}
	}
}
