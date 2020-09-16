package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	/*
		ioiutil
			ReadFile()
			WriteFile()
			ReadDir()
	*/
	fileName := "abc.txt"
	//ReadFile读取文件到一个切片
	data, err := ioutil.ReadFile(fileName)
	fmt.Println(err)
	fmt.Println(string(data))

	//WriteFile将一个切片写出到文件(会清空文件)
	fileName2 := "cba.txt"
	err = ioutil.WriteFile(fileName2, []byte("hello"), os.ModePerm)
	fmt.Println(err)

	//ReadAll 读取一个Reader对象读成切片
	reader := strings.NewReader("test")
	data, err = ioutil.ReadAll(reader)
	fmt.Println(err)
	fmt.Println(data)

	//ReadDir 读取一个文件夹下所有文件  返回fileInfo类型
	dirName := "D:/gopath/src/advanceProject"
	fileList,err :=ioutil.ReadDir(dirName)
	fmt.Println(err)
	for i:=0;i<len(fileList);i++{
		fmt.Println(i," ",fileList[i].Name())
	}


	//TempDir 临时目录
	dir,err:=ioutil.TempDir("D:/gopath/src/advanceProject","tmp")
	fmt.Println(err)
	fmt.Println(dir)
	defer os.RemoveAll(dir)

	//TempFile 创建临时文件 返回fileInfo
	file,err:=ioutil.TempFile(dir,"tmp")
	fmt.Println(err)
	fmt.Println(file.Name())
	//defer os.Remove(file.Name())
	defer file.Close()
}
