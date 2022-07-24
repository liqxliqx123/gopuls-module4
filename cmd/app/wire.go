//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"module4/conf"
	"module4/data"
)

func Initialize() *data.MyDB {
	wire.Build(conf.NewConfiure, data.NewMyDB)
	return &data.MyDB{}
}
