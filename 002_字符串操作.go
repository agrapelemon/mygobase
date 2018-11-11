package main

import (
	"fmt"
	"strings"
)

func main() {
	if true {
		fmt.Println("--1.字符串相关操作---------------")

		fmt.Println("----1.1.修改字符串，只能是英文，中文报错---------------")
		a1 := "a哈哈"
		temp1 := []byte(a1) // 转换类型
		temp1[0] = 'b'
		b1 := string(temp1[:]) //转回字符串
		fmt.Println(b1)

		fmt.Println("----1.2.字符串连接，两个元素都只能是字符串---------------")
		a2 := "aa"
		temp2 := "bb"
		b2 := a2 + temp2
		fmt.Println(b2)

		fmt.Println("----1.3.是否包含某个字符串---------------")
		b3 := strings.Contains("是否包含某个字符串", "字符串")
		fmt.Println(b3)

		fmt.Println("----1.4.字符串分割---------------")
		b4 := strings.Split("是否包含某个字符串", "某个")
		fmt.Printf("%T %+v\n", b4, b4)

		fmt.Println("----1.5.字符串截取，只能是英文---------------")
		a5 := "asf是否包含某个字符串"
		b5 := a5[2:]
		fmt.Printf("%T %+v\n", b5, b5)

		fmt.Println("----1.6.字符串截取，中文---------------")
		a6 := "asf是否包含某个字符串"
		temp6 := []rune(a6)
		b6 := string(temp6[4:])
		fmt.Printf("%T %+v\n", b6, b6)

		fmt.Println("----1.7.修改字符串---------------")
		a7 := "asf是否包含某个字符串"
		temp7 := []rune(a7)
		temp7[0] = '张'
		b7 := string(temp7)
		fmt.Printf("%T %+v\n", b7, b7)

		fmt.Println("----1.8.判断字符串是否相等---------------")
		a8 := "asf是否包含某个字符串"
		temp8 := "asf是否包含某个字符串"
		b8 := a8 == temp8
		fmt.Printf("%T %+v\n", b8, b8)

	}
}
