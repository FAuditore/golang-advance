package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func initMysql() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	return db, err
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	db, err := initMysql()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("数据库连接成功")
	defer db.Close()

	v1 := router.Group("/user")
	{
		v1.POST("/login", func(c *gin.Context) {

		})

		v1.GET("/list", func(c *gin.Context) {
			users := []User{}
			db.Debug().Find(&users)
			c.JSON(200,users)
		})

		v1.GET("/createtable", func(c *gin.Context) {
			db.Debug().AutoMigrate(&User{})
			c.String(200, fmt.Sprintf("%v", db))
		})

		v1.POST("/register", func(c *gin.Context) {
			account := c.PostForm("account")
			password := c.PostForm("password")
			user := User{
				Account:  account,
				Password: password,
			}
			db.Debug().Create(&user)
			c.String(200, "注册成功")
		})

		v1.PUT("/modify", func(c *gin.Context) {
			user := User{
				Account:  "abc",
				Password: "123",
			}
			db.Debug().Model(&User{}).Where("account=?", "foo").Updates(user)
		})

		v1.DELETE("/delete", func(c *gin.Context) {
			db.Debug().Delete(User{}, "id=?", 1)
		})

	}
	router.Run(":8090")
}
