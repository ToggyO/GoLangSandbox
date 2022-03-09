package server

import (
	"fmt"
	"net"
)

type App struct{}

func (a *App) Start(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go HandleConnection(conn)
	}
}
