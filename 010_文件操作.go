package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	if false {
		fmt.Println("--1.创建文件，有文件覆盖。写文件内容---------------")
		file, err := os.Create("010/1.txt") //覆盖内容
		//file, err := os.OpenFile("010/test/1.txt", os.O_WRONLY|os.O_APPEND, 0666) //追加内容
		if err != nil {
			fmt.Println("创建文件失败！", err)
			return
		}
		defer file.Close()
		fmt.Println("创建文件成功！", file.Name())
		file.WriteString("写字符串\r\n")
		file.Write([]byte("写二进制\r\n"))
	}
	if false {
		fmt.Println("--2.删除文件---------------")
		err := os.Remove("010/1.txt")
		if err != nil {
			fmt.Println("删除文件失败！", err)
			return
		}
		fmt.Println("删除文件成功！")
	}
	if false {
		fmt.Println("--3.打开文件。读文件内容。---------------")
		//file, err := os.OpenFile("./1.txt", os.O_RDWR|os.O_CREATE, 0766)
		file, err := os.Open("./010/1.txt")
		if err != nil {
			fmt.Println("打开文件失败！", err)
			return
		}
		defer file.Close()
		fmt.Println("打开文件成功！", file.Name())
		//创建byte的slice用于接收文件读取数据
		buf := make([]byte, 1024)
		//循环读取
		for {
			//Read函数会改变文件当前偏移量
			len, _ := file.Read(buf)
			//读取字节数为0时跳出循环
			if len == 0 {
				break
			}
			fmt.Println(string(buf[:len]))
		}
		bytes2, err2 := ioutil.ReadFile("./010/1.txt")
		if err2 != nil {
			fmt.Println("读取文件失败！", err2)
		}
		fmt.Println("读取全部文件内容！#", string(bytes2), "#")
	}
	if false {
		fmt.Println("--4.文件或目录是否存在。是否是文件---------------")
		fi, err := os.Stat("./010/1.txt")
		if err != nil {
			fmt.Println("--文件或目录不存在。---------------")
			return
		}
		fmt.Println("--文件或目录存在。---------------")
		fmt.Println("是否是文件：", !fi.IsDir())
	}
	if false {
		fmt.Println("--5.创建目录。---------------")
		err2 := os.Mkdir("010/test", os.ModePerm)
		if err2 != nil {
			fmt.Println("创建文件夹失败！", err2)
			return
		}
		file, err := os.Create("010/test/1.txt") //覆盖内容
		//file, err := os.OpenFile("010/test/1.txt", os.O_WRONLY|os.O_APPEND, 0666) //追加内容
		if err != nil {
			fmt.Println("创建文件失败！", err)
			return
		}
		defer file.Close()

	}
	if false {
		fmt.Println("--6.删除目录（或文件）。---------------")
		err := os.RemoveAll("010/1.txt")
		if err != nil {
			fmt.Println("删除目录失败！", err)
			return
		}
		fmt.Println("创建目录成功！")
	}
	if false {
		fmt.Println("--7.文件或目录重命名。可以移动文件夹---------------")
		err := os.Rename("010/test2", "010/test3")
		//err := os.Rename("029", "010/test")
		if err != nil {
			fmt.Println("文件或目录重命名失败！", err)
			return
		}
		fmt.Println("文件或目录重命名成功！")
	}
	if false {
		fmt.Println("--8.文件复制---------------")
		_, err := copy("010/2.txt", "010/3.txt")
		//err := os.Rename("029", "010/test")
		if err != nil {
			fmt.Println("文件复制失败！", err)
			return
		}
		fmt.Println("文件复制成功！")
	}
	if true {
		fmt.Println("--9.文件夹复制，未做---------------")
		//copyDir("010", "0101")
		fmt.Println("文件夹复制成功！")
	}
	if true {
		fmt.Println("--10.获取文件夹里的文件和文件夹---------------")
		rd, err := ioutil.ReadDir("010")
		if err != nil {
			return
		}
		for _, fi := range rd {
			fmt.Println(fi.Name())
		}
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}
	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
