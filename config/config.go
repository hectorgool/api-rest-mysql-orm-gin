package config

import (
    "fmt"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/hectorgool/gin2/model"
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
    //defer Dbh.Close()

    Dbh.AutoMigrate(&model.Person{})

}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            fmt.Println("OPTIONS")
            c.AbortWithStatus(200)
        } else {
            c.Next()
        }
    }
}
