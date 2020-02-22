package initializer

import (
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var (
	DB *xorm.Engine
)

func InitXORM() {
	DB, err := xorm.NewEngine(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, "postgres", "password", "reputation"),
	)
	if err != nil {
		panic(err)
	}
	DB.ShowSQL()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connect postgresql success")
}
