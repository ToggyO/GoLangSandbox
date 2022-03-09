package server

import (
	"fmt"
	"http.app/cmd/server/data"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, 1024*2)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}

		source := string(input[:n])
		target, ok := server.Translation[source]
		if !ok {
			target = "undefined"
		}

		fmt.Println(source, " - ", target)

		conn.Write([]byte(target))
	}
}
