package initializer

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"xorm.io/xorm"

	"github.com/mingchoi/hello-micro-service/service-entity-reputation/app"
)

func InitXORM() {
	time.Sleep(5 * time.Second)

	dataSource := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.Get("db.host"),
		viper.Get("db.port"),
		viper.Get("db.user"),
		viper.Get("db.password"),
		viper.Get("db.dbname"),
	)

	var err error
	app.DB, err = xorm.NewEngine("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	app.DB.ShowSQL()

	err = app.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to postgres")
}
