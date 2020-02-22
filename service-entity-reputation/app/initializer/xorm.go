package initializer

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var (
	DB *xorm.Engine
)

func InitXORM() {
	time.Sleep(5 * time.Second)
	DB, err := xorm.NewEngine(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "postgres", 5432, "postgres", "password", "reputation"),
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
