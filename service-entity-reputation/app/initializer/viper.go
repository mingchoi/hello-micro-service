package initializer

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

var defaultConfig = []byte(`
db:
  host: localhost
  port: 5432
  user: postgres
  password: password
  dbname: reputation
app:
  grpc:
    port: 50051
`)

func InitViper() {
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("ENTITY_REPUTATION")
}
