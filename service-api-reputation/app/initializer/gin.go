package initializer

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mingchoi/hello-micro-service/service-api-reputation/app/handler"
	"github.com/spf13/viper"
)

func InitGin() {
	port := ":" + strconv.Itoa(viper.Get("app.port").(int))
	r := gin.Default()
	r.GET("/", handler.Hello)
	r.GET("/call", handler.Handle)
	r.Run(port)
}
