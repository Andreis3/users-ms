package main

import (
	_interface "github.com/andreis3/users-ms/src/interface"
	userInterface "github.com/andreis3/users-ms/src/interface/http/presentation/user"
)

var server userInterface.UserRouter

func main() {
	server := _interface.NewServer(server)
	server.Start()
}
