package main

import (
	"ego/internal/server"
)

func main() {
	server.CreateServer().Start(":8081")
}
