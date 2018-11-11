package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/timespacegroup/go-utils"
	"os/exec"
)

func main() {
	if true {
		// 创建
		u1, _ := uuid.NewV4()
		fmt.Println("--1.生成uuid---------------")
		fmt.Printf("UUIDv4: %s\n", u1)

		fmt.Println("--2.字符串转uuid---------------")
		u2, _ := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
		fmt.Println(u2)
		fmt.Println("--3.uuid转字符串---------------")
		s3 := u1.String()
		fmt.Println(s3)
		fmt.Println("--4.uuid转byte数组---------------")
		s4 := u1.Bytes()
		fmt.Println(s4)
		fmt.Println("--5.byte数组转uuid---------------")
		s5, _ := uuid.FromBytes(s4)
		fmt.Println(s5)
	}
	if true {
		fmt.Println("--1.生成guid，只有字符串，没有byte数组，慎用---------------")
		var u1 string
		u1 = tsgutils.GUID()
		fmt.Println(u1)
		fmt.Println("--2.生成uuid，只有字符串，没有byte数组，慎用---------------")
		u1 = tsgutils.UUID()
		fmt.Println(u1)
	}
	if true {
		fmt.Println("--1.生成uuid,windows下生成失败---------------")
		out, _ := exec.Command("uuidgen").Output()
		//if err != nil {
		//	log.Fatal(err)
		//}
		fmt.Printf("%s", out)
	}
}
