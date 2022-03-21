package main

import "http.app/cmd/server/net"

func main() {
	var app server.App

	app.Start("8080")
}
