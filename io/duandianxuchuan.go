package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	srcFileName := "D:/gopath/src/advanceProject/abc.txt"
	destFileName := srcFileName[strings.LastIndex(srcFileName, "/")+1:]
	fmt.Println(destFileName)
	tempFileName := destFileName + "temp.txt"
	fmt.Println(tempFileName)

	srcFile, err := os.Open(srcFileName)
	HandleError(err)
	destFile, err := os.OpenFile(destFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandleError(err)
	tempFile, err := os.OpenFile(tempFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	HandleError(err)
	defer srcFile.Close()
	defer destFile.Close()

	tempFile.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, _ := tempFile.Read(bs)
	countStr := string(bs[:n1])
	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println(count)

	srcFile.Seek(count, io.SeekStart)
	destFile.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024)
	n2 := -1 //读取数据量
	n3 := -1 //写出数据量
	total := int(count)
	for {
		n2, err = srcFile.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕")
			tempFile.Close()
			os.Remove(tempFileName)
			break
		}
		n3, err = destFile.Write(data[:n2])
		total += n3
		//将复制总量存储到文件中
		tempFile.Seek(0, io.SeekStart)
		tempFile.WriteString(strconv.Itoa(total))
	}
}
