package main

import (
	"fmt"
	"time"
)

func main() {
	if true {
		fmt.Println("--1.获取当前时间---------------")
		var a1 time.Time = time.Now()
		fmt.Println(a1)
		fmt.Println("--2.日期时间转换成字符串---------------")
		var a2 string = a1.Format("2006-01-02 15:04:05")
		fmt.Println(a2)
		fmt.Println("--3.日期时间转换成时间戳（秒）---------------")
		var a3 int64 = a1.Unix()
		fmt.Println(a3)
		fmt.Println("--4.时间戳（秒）转换成日期时间---------------")
		var a4 time.Time = time.Unix(a3, 0)
		fmt.Println(a4)
		//fmt.Println(time.Now().UnixNano()) //获取当前纳秒
		fmt.Println("--5.字符串转换成日期时间---------------")
		a5, _ := time.Parse("01/02/2006", "02/08/2015")
		fmt.Println(a5)

		fmt.Println("--6.字符串转换成时间戳（秒）---------------")
		s6 := "2016-01-02 15:04:05"
		temp6, _ := time.Parse("2006-01-02 15:04:05", s6)
		a6 := temp6.Unix()
		fmt.Println(a6)

		fmt.Println("--7.时间戳（秒）转换成字符串---------------")
		var u7 int64 = a6
		a7 := time.Unix(u7, 0).Format("2006-01-02 15:04:05")
		fmt.Println(a7)

	}
	if false {
		datetime := "2015-01-01 00:00:00" //待转化为时间戳的字符串

		//日期转化为时间戳
		timeLayout := "2006-01-02 15:04:05"  //转化所需模板
		loc, _ := time.LoadLocation("Local") //获取时区
		tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
		timestamp := tmp.Unix() //转化为时间戳 类型是int64
		fmt.Println(timestamp)

		//时间戳转化为日期
		datetime = time.Unix(timestamp, 0).Format(timeLayout)
		fmt.Println(datetime)
	}
}
