package main

import (
	"fmt"
	"github.com/MrDienns/gin-example/internal/web"
)

func main() {
	server := web.NewServer()

	err := server.Start()
	if err != nil {
		fmt.Println("error starting webserver: " + err.Error())
	}
}
