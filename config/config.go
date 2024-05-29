package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	Viper() *viper.Viper
}
