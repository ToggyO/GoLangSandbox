package main

import dotnet "http.app/cmd/client/net"

func main() {
	var client dotnet.TcpClient

	client.Start("8080")
}
