package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"unsafe"
)

type Person struct {
	Id   int
	Name string
}
type Student struct {
	Person
	Name2 string
}
type (
	文本 string
)

//func nameof(iii interface{}) {
//
//	fmt.Println("nameof测试")
//}

//func nameof(iii Type) {
//
//	fmt.Println("nameof测试")
//}
func main() {

	//aaa:=new Person{}
	//var aaa int = 0
	//nameof(Person)
	if true {
		fmt.Println("--1.类型自动推导---------------")
		var a1 = 311111111111.11111
		var a2 = "哈哈"
		var a3 = true
		var a4 = 23
		var a5 = complex(1, 2)
		var a6 = 'a'
		var a7 = new(Person)              //结构体指针
		var a8 = Person{}                 //结构体
		var a9 = func() {}                //函数
		var a10 = [...]int{1, 2, 3, 4, 5} //数组
		var a11 = []int{1, 2, 3, 4, 5}    //切片
		var a12 = make([]int, 5)          //切片
		var a13 = []interface{}{1, 2, 3}
		var a14 = map[string]string{"France": "Paris", "Italy": "Rome"}

		fmt.Printf("%T\n", a1)
		fmt.Printf("%T\n", a2)
		fmt.Printf("%T\n", a3)
		fmt.Printf("%T\n", a4)
		fmt.Printf("%T\n", a5)
		fmt.Printf("%T\n", a6)
		fmt.Printf("%T\n", a7)
		fmt.Printf("%T\n", a8)
		fmt.Printf("%T\n", a9)
		fmt.Printf("%T\n", a10)
		fmt.Printf("%T\n", a11)
		fmt.Printf("%T\n", a12)
		fmt.Printf("%T\n", a13)
		fmt.Printf("%T\n", a14)
	}
	if true {
		fmt.Println("--2.类型自动转换---------------")

		var a1 int = 65
		b1 := string(a1)
		fmt.Printf("整型转字符串%T %+v\n", b1, b1)

		var a2 string = "88"
		b2, _ := strconv.ParseInt(a2, 10, 32)
		fmt.Printf("字符串转整型%T %+v\n", b2, b2)

		var a3 string = "88"
		b3, _ := strconv.Atoi(a3)
		fmt.Printf("字符串转整型%T %+v\n", b3, b3)

		var a4 int = 65
		b4 := fmt.Sprintf("test%d", a4)
		fmt.Printf("整型转字符串%T %+v\n", b4, b4)

		var a5 string = "hd尕的三法司"
		var b5 interface{} = a5
		fmt.Printf("字符串转空接口%T %+v\n", b5, b5)

		var a6 interface{} = "hd尕的三法司"
		var b6 string
		b6 = a6.(string)
		//b6 = string(a6)
		fmt.Printf("接口转字符串%T %+v\n", b6, b6)

		var a7 string = "hd尕的三法司"
		var b7 文本
		//b7=a7
		b7 = 文本(a7)
		fmt.Printf("字符串转文本%T %+v\n", b7, b7)

		var a8 文本 = "hd尕的三法司"
		var b8 文本
		b8 = a8
		fmt.Printf("字符串转文本%T %+v\n", b8, b8)

		var a9 int = 123
		var b9 float64
		b9 = float64(a9)
		fmt.Printf("整型转浮点型%T %+v\n", b9, b9)

		var a10 float64 = 123
		var b10 int
		b10 = int(a10)
		fmt.Printf("浮点型转整型%T %+v\n", b10, b10)

		//var a11=Student{}//转换失败，当前类不能转父类
		//var b11 Person
		//b11=Person(a11)
		//fmt.Printf("当前类转父类%T %+v\n", b11,b11)
		a12 := 234
		b12 := fmt.Sprintf("%daaa", a12)
		fmt.Printf("整型转字符串： %s\n", b12)
	}
	if true {
		fmt.Println("--3.unsafe---------------")

		s := int16(0x1234)
		b := int8(s)
		fmt.Println("int16字节大小为", unsafe.Sizeof(s)) //结果为2
		if 0x34 == b {
			fmt.Println("little endian")
		} else {
			fmt.Println("big endian")
		}

	}
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

/*
 输入输出。这个分类包括二进制以及文本格式在屏幕、键盘、文件以及其他设备上的输 入输出等，比如二进制文件的读写。对应于此分类的包有bufio、fmt、io、log和flag 等，其中flag用于处理命令行参数。
文本处理。这个分类包括字符串和文本内容的处理，比如字符编码转换等。对应于此分 类的包有encoding、bytes、strings、strconv、text、mime、unicode、regexp、 index和path等。其中path用于处理路径字符串。
网络。这个分类包括开发网络程序所需要的包，比如Socket编程和网站开发等。对应于此 分类的包有：net、http和expvar等。
系统。这个分类包含对系统功能的封装，比如对操作系统的交互以及原子性操作等。对 应于此分类的包有os、syscall、sync、time和unsafe等。
数据结构与算法。对应于此分类的包有math、sort、container、crypto、hash、 archive、compress和image等。因为image包里提供的图像编解码都是算法，所以也 归入此类。
运行时。对应于此分类的包有：runtime、reflect和go等。

A.1.1 常用包介绍 本节我们介绍Go语言标准库里使用频率相对较高的一些包。熟悉了这些包后，使用Go语言 开发一些常规的程序将会事半功倍。
fmt。它实现了格式化的输入输出操作，其中的fmt.Printf()和fmt.Println()是开 发者使用最为频繁的函数。
io。它实现了一系列非平台相关的IO相关接口和实现，比如提供了对os中系统相关的IO 功能的封装。我们在进行流式读写（比如读写文件）时，通常会用到该包。
bufio。它在io的基础上提供了缓存功能。在具备了缓存功能后，bufio可以比较方便地 提供ReadLine之类的操作。
strconv。本包提供字符串与基本数据类型互转的能力。 
os。本包提供了对操作系统功能的非平台相关访问接口。接口为Unix风格。提供的功能 包括文件操作、进程管理、信号和用户账号等。
sync。它提供了基本的同步原语。在多个goroutine访问共享资源的时候，需要使用sync 中提供的锁机制。
flag。它提供命令行参数的规则定义和传入参数解析的功能。绝大部分的命令行程序都 需要用到这个包。
encoding/json。JSON目前广泛用做网络程序中的通信格式。本包提供了对JSON的基 本支持，比如从一个对象序列化为JSON字符串，或者从JSON字符串反序列化出一个具体 的对象等。 
http。它是一个强大而易用的包，也是Golang语言是一门“互联网语言”的最好佐证。通 过http包，只需要数行代码，即可实现一个爬虫或者一个Web服务器，这在传统语言中 是无法想象的。
*/
