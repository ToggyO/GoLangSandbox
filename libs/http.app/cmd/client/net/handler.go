package client

import (
	"fmt"
	"net"
	"time"
)

func Handle(conn net.Conn) {
	for {
		var source string
		fmt.Println("Ввуедите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}

		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Перевод:")
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))

		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}

			fmt.Println(string(buff[:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
	}
}
