package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "bufio"
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//创建Reader对象  缓冲区默认大小4096
	//b1:=bufio.NewReader(file)

	//1.读取到p切片中 Read
	//p :=make([]byte,1024)
	//n1,err := b1.Read(p)
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	//2.ReadLine()
	//data,flag,err:=b1.ReadLine()
	//if err!=nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(data))
	//fmt.Println(flag)

	//3.ReadString
	//s1, err := b1.ReadString('\n')
	//fmt.Println(s1)
	//for {
	//	s1, err = b1.ReadString('\n')
	//	if err==io.EOF{
	//		break
	//	}
	//	fmt.Println(s1)
	//}

	//4.ReadBytes
	//data,err:=b1.ReadBytes('\n')
	//fmt.Println(string(data))

	//Scanner
	s2 := ""
	//fmt.Scanln只读到空格
	//fmt.Scanln(&s2)
	//fmt.Println(s2)
	b2 := bufio.NewReader(os.Stdin)
	s2, err = b2.ReadString('\n')
	fmt.Print(s2)

	//写入到缓冲区中
	file.Seek(0, io.SeekEnd)
	writer := bufio.NewWriter(file)
	n, err := writer.WriteString(s2)
	fmt.Println(err)
	fmt.Println(n)

	//将缓冲区内容刷新
	err = writer.Flush()
	fmt.Println(err)
}
