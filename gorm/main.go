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
	IsOne     bool
	IsTwo     bool `gorm:"type:bool"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}

type BoardToken struct {
	ID          int64
	AppId       int
	Type        string
	AeolusAppId string
	SecretKey   string
	AccessToken string
	Creator     string
	Status      bool
	CreatedTime time.Time
	UpdateTime  time.Time
}

func (BoardToken) TableName() string {
	return "board_token"
}

func main() {
	//dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       dsn,
	//	SkipInitializeWithVersion: true,
	//}),
	//&gorm.Config{
	//	DisableAutomaticPing: false,
	//	DryRun: true})
	//
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}

	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&BoardToken{})

	a := &BoardToken{
		AppId:       3241,
		Type:        "431",
		AeolusAppId: "4321",
		SecretKey:   "4321",
		AccessToken: "4312",
		Creator:     "432141",
		Status:      true,
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
	}
	fmt.Println(db.Create(a).Error)
}
