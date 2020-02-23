package main

import (
	"github.com/mingchoi/hello-micro-service/service-api-reputation/app/initializer"
)

func main() {
	initializer.InitViper()

	conn := initializer.InitGRPC()
	defer conn.Close()

	initializer.InitGin()
}
