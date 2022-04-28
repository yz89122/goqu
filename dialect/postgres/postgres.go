package postgres

import (
	"github.com/yz89122/goqu/v10"
)

func DialectOptions() *goqu.SQLDialectOptions {
	do := goqu.DefaultDialectOptions()
	do.PlaceHolderFragment = []byte("$")
	do.IncludePlaceholderNum = true
	return do
}

func init() {
	goqu.RegisterDialect("postgres", DialectOptions())
}
