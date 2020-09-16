package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	fileInfo, err := os.Stat("abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode()) //权限
	fmt.Println(filepath.Abs("test.txt"))
	fmt.Println(path.Join("D:/gopath/src/", ".."))

	//MkdirAll 递归创建
	err = os.MkdirAll("D:/gopath/src/advanceProject/a/b/c/d", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	//create创建文件
	file, err := os.Create("abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	fileInfo2, _ := file.Stat()
	fmt.Println(fileInfo2.ModTime())
	file.Close()

	//Open()只读方式打开
	//打开文件，OpenFile(filename,mode(打开方式),perm(权限))
	file2, err := os.OpenFile("abc.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file2)
	file2.Close()

	file3, _ := os.Create("test.txt")
	file3.Close()
	//删除
	err = os.Remove("test.txt")
	if err != nil {
		fmt.Println(err)
	}

}
