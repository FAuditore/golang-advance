package main

import (
	. "fmt"
	"time"
)

func main() {
	t1 := time.Now()
	Println(t1)

	//日期格式化模板 Mon Jan(1) 2 15(03代表12小时制，15代表24小时制):04:05 MST 2006
	s1 := t1.Format("Mon 2006-01-02 15:04:05 ")
	Println(s1)
	hour, min, sec := t1.Clock()
	Println(hour, min, sec)

	//string->time
	s2 := "2020年7月25日"
	t2, err := time.Parse("2006年1月2日", s2)
	if err != nil {
		println(err)
	}
	Println("t2: ", t2)

	Println(t1.Weekday())
	Println(t1.YearDay())

}
