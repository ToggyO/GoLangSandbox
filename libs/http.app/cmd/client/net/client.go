package client

import (
	"fmt"
	"net"
)

type TcpClient struct{}

func (t TcpClient) Start(port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	Handle(conn)
}
