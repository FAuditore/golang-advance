package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"column:user_name;type:varchar(20)"`
	Age       int    `gorm:"column:user_age"`
	Sex       int    `gorm:"column:sex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}

func main() {
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	//_ = db.AutoMigrate(&User{
	//	ID:        1,
	//	Name:      "abc",
	//	Age:       1,
	//	Sex:       1,
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//})

	fmt.Println(db.Create(&User{
		ID:        1,
		Name:      "liubo",
		Age:       10,
		Sex:       0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error)


	statement := db.Session(&gorm.Session{
		DryRun:  true,
	}).First(&User{}).Statement
	fmt.Println(statement.SQL.String())

}
