package initializer

import (
	"bytes"
	"strings"

	"github.com/spf13/viper"
)

var defaultConfig = []byte(`
app:
  port: 8000
`)

func InitViper() {
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("API_REPUTATION")
}
