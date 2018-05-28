package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hectorgool/api-rest-mysql-orm-gin/model"
	"github.com/jinzhu/gorm"
	"os"
)

var Dbh *gorm.DB
var err error

func init() {

	dbParams := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE_NAME"),
		os.Getenv("MYSQL_CHARSET"),
		os.Getenv("MYSQL_PARSETIME"),
		os.Getenv("MYSQL_LOC"))

	Dbh, _ = gorm.Open("mysql", dbParams)
	if err != nil {
		fmt.Println(err)
	}

	Dbh.AutoMigrate(&model.Person{})

}
