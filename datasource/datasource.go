package datasource

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// ConnRDB is ...
var ConnRDB *gorm.DB

// type DBConnection struct {
// 	RDB *gorm.DB
// }

// CreateRDB is ...
func CreateRDB() {
	var err error
	ConnRDB, err = gorm.Open("mysql", "root:1234@tcp(localhost:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal("ERROR RDB ", err)
	}
	ConnRDB.DB().SetMaxIdleConns(10)
	ConnRDB.DB().SetMaxOpenConns(100)
	ConnRDB.DB().SetConnMaxLifetime(3 * time.Minute)
	//ConnRDB.SetLogger(log.New(os.Stdout, "\r\n", 0))
	ConnRDB.LogMode(true)
	ConnRDB.SingularTable(true)
}
