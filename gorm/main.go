package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"column:user_name;type:varchar(20)"`
	Age  int    `gorm:"column:user_age"`
	Sex  int    `gorm:"column:sex"`
}

func (User) TableName() string {
	return "user"
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//db.Debug().DropTable(&User{})

	//自动添加缺失field 不会删除原field
	db.Debug().AutoMigrate(&User{})

	//创建表
	//db.Table("user").CreateTable(&User{})

	user := User{Name: "liubo", Age: 1000, Sex: 1}
	fmt.Println(db.NewRecord(&user))


	//插入
	//db.Debug().Create(&user)

	//查询
	//.First 第一个
	db.Debug().First(&user)
	//db.First(&user,2)
	fmt.Println(user)
	//select Select 选择读取字段
	//db.Debug().Select("user_name,user_age").First(&user)
	users:=[]User{}
	//unscoped 包含软删除的结果
	db.Unscoped().Where("user_age<>1").Find(&users)
	fmt.Println(users)

	//更新
	user = User{}
	//Save修改所有字段
	//user.Name="abc"
	//db.Debug().Save(&user)

	//Update
	//db.Debug().Model(&user).Update("user_name","哈哈")

	//Updates
	//db.Debug().Model(&user).Updates(user)
	//db.Debug().Model(&user).Where("id=?",user.ID).Updates(map[string]interface{}{"user_name":"hello","user_age":18,"sex":0})


	//删除 若无ID值，则全部删除（软删除）
	//db.Debug().Delete(&user)
	//db.Debug().Where("user_name=?","liubo").Delete(&user)
	//db.Debug().Delete(&User{},"user_age=?",1)
	//物理删除
	//db.Debug().Unscoped().Delete(&User{},"sex=?",0)
}
