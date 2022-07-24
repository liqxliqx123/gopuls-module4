package conf

import "github.com/spf13/viper"

var vip *viper.Viper

type Configure struct{}

func NewConfiure() *viper.Viper {
	if vip == nil {
		vip = viper.New()
	}
	return vip
}
