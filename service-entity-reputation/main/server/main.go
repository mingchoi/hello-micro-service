package main

import (
	"github.com/mingchoi/hello-micro-service/service-entity-reputation/app/initializer"
)

func main() {
	initializer.InitViper()

	initializer.InitXORM()

	initializer.InitGRPC()
}
