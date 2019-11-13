package middleware

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/senseoki/iris_ex/datasource"
)

// RdbTX is ...
func RdbTX(ctx iris.Context) {

	var isTx bool
	tx := datasource.ConnRDB

	switch ctx.Method() {
	case "POST", "PUT", "DELETE":
		isTx = true
		tx = tx.Begin()
		ctx.Values().Set("RDBTX", tx)
	default:
		ctx.Values().Set("RDBTX", tx)
	}

	defer func(isTx bool) {
		if isTx {
			if ctx.GetStatusCode() >= 500 {
				log.Println("RDB TX Rollback")
				tx.Rollback()
			} else {
				log.Println("RDB TX Commit")
				tx.Commit()
			}
		}
	}(isTx)

	ctx.Next()
}
