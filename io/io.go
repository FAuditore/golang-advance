package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	/*
		Reader接口：
			Read(p []byte)(n int, error)
	*/
	//1.打开文件
	fileName := "abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	//3.关闭文件
	defer file.Close()

	//2.读取数据

	//创建切片存储读取数据
	bs := make([]byte, 4, 4)
	/*
		//n 读取到的字节数
		n, err := file.Read(bs)
		fmt.Println(err)
		fmt.Println(n)
		fmt.Println(string(bs))

		n,err = file.Read(bs)
		fmt.Println(err)
		fmt.Println(n)
		fmt.Println(string(bs))

		n,err = file.Read(bs)
		fmt.Println(err)
		fmt.Println(n)
		fmt.Println(bs)
	*/

	n := -1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读完了")
			break
		}
		fmt.Print(string(bs[:n]))
	}

	//写文件
	fileName2 := "write.txt"
	//1.打开文件
	file2, err2 := os.OpenFile(fileName2, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	HandleError(err2)
	defer file2.Close()

	//2.写出数据
	bs2 := []byte("write!!!")
	n2, err2 := file2.Write(bs2)
	HandleError(err2)
	fmt.Println(n2)
	//写字符串
	n2, err2 = file2.WriteString("哈哈")
	HandleError(err2)
	fmt.Println(n2)

	//复制文件
	total3, err3 := copyFile("write.txt", "copy.txt")
	HandleError(err3)
	fmt.Println(total3)

	//io.Copy(dest,src)
	total4, err4 := copyFile2("abc.txt", "cba.txt")
	HandleError(err4)
	fmt.Println(total4)

	//ioutil.ReadFile  ioutil.WriteFile 一次性全部读取到一个切片中
	total5, err5 := copyFile3("abc.txt", "aabbcc.txt")
	HandleError(err5)
	fmt.Println(total5)

	//断点续传
	file6, err6 := os.OpenFile("aabbcc.txt", os.O_RDWR, os.ModePerm)
	HandleError(err6)
	defer file6.Close()

	//Seek(offset,whence)   abcdefg123
	bs6 := []byte{0}
	file6.Read(bs6)
	fmt.Println(string(bs6)) //a

	//从SeekStart开始偏移4个字节开始读，即从第五个字节读取
	//SeekCurrent 当前位置  SeekEnd 结尾
	file6.Seek(4, io.SeekStart)
	file6.Read(bs6)
	fmt.Println(string(bs6)) //e

	file6.Seek(2, io.SeekCurrent)
	file6.Read(bs6)
	fmt.Println(string(bs6)) //1
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copyFile(srcFile, destFile string) (int, error) {
	//打开
	sfile, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	dfile, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	//关闭
	defer sfile.Close()
	defer dfile.Close()

	//读写
	bs := make([]byte, 1024, 1024)
	n := -1 //读取数据量
	total := 0
	for {
		n, err = sfile.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			return total, err
		}
		total += n
		dfile.Write(bs[:n])
	}
	return total, nil
}

//默认32*1024的切片
func copyFile2(srcFile, destFile string) (int64, error) {
	sfile, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	dfile, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer sfile.Close()
	defer dfile.Close()
	return io.Copy(dfile, sfile)
}

//全部读取到内存，再拷贝
func copyFile3(srcFile, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}
	return len(bs), err
}
