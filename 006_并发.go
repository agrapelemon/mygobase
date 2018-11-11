package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}
func main() {
	if false {
		fmt.Println("--1.Gosched的使用---------------")
		go func() {
			for i := 0; i < 5; i++ {
				fmt.Println("go")
			}
		}()
		for i := 0; i < 2; i++ {
			//让出时间片，先让别的协程执行，它执行完，再回来执行
			runtime.Gosched()
			fmt.Println("hello")
		}
	}
	if false {
		fmt.Println("--2.Goexit的使用---------------")
		go func() {
			for i := 0; i < 5; i++ {
				fmt.Println("go")
			}
		}()
		for i := 0; i < 2; i++ {
			//终止所在协程
			runtime.Goexit()
			fmt.Println("hello")
		}
	}
	if false {
		fmt.Println("--3.GOMAXPROCS的使用---------------")
		n := runtime.GOMAXPROCS(40)
		fmt.Println("n = ", n)
		for {
			go fmt.Print(1)
			fmt.Print(0)
		}
	}
	if false {
		fmt.Println("--4.死锁---------------")
		var ch = make(chan int)
		ch <- 666
		go func() {
			for {
				<-ch
				fmt.Print("h")
				time.Sleep(time.Second)
				fmt.Print("e")
				time.Sleep(time.Second)
				fmt.Print("l")
				time.Sleep(time.Second)
				fmt.Print("l")
				time.Sleep(time.Second)
				fmt.Println("o")
				time.Sleep(time.Second)
				ch <- 666
			}

		}()
		go func() {
			for {
				<-ch
				fmt.Print("w")
				time.Sleep(time.Second)
				fmt.Print("o")
				time.Sleep(time.Second)
				fmt.Print("r")
				time.Sleep(time.Second)
				fmt.Print("l")
				time.Sleep(time.Second)
				fmt.Println("d")
				time.Sleep(time.Second)
				ch <- 666
			}
		}()

		time.Sleep(60 * time.Second)

	}
	if false {
		fmt.Println("--5.无缓存通道，解决同步问题。写入数据，当前协程暂停---------------")
		var ch = make(chan int, 0)
		//go func() {
		//	ch <- 666 //必须另外开协程写数据，在主协程中写数据，会死锁
		//}()

		go func() {
			for {
				<-ch
				fmt.Print("h")
				time.Sleep(time.Second)
				fmt.Print("e")
				time.Sleep(time.Second)
				fmt.Print("l")
				time.Sleep(time.Second)
				fmt.Print("l")
				time.Sleep(time.Second)
				fmt.Println("o")
				time.Sleep(time.Second)
				ch <- 666 //写数据
			}

		}()
		go func() {
			for {
				<-ch //读数据
				fmt.Print("w")
				time.Sleep(time.Second)
				fmt.Print("o")
				time.Sleep(time.Second)
				fmt.Print("r")
				time.Sleep(time.Second)
				fmt.Print("l")
				time.Sleep(time.Second)
				fmt.Println("d")
				time.Sleep(time.Second)
				ch <- 666 //写数据
			}
		}()
		ch <- 666 //有协程，会死锁
		time.Sleep(60 * time.Second)

	}
	if false {
		fmt.Println("--6.关闭通道，不可写(写的时候抛出panic)，可读（读的时候是0）---------------")
		var ch = make(chan int, 1)
		//go func() {
		//	ch <- 666 //必须另外开协程写数据，在主协程中写数据，会死锁
		//}()

		go func() {
			for {
				iii := <-ch
				fmt.Println("iii = ", iii)
				fmt.Print("h")
				//time.Sleep(time.Second)
				fmt.Print("e")
				//time.Sleep(time.Second)
				fmt.Print("l")
				//time.Sleep(time.Second)
				fmt.Print("l")
				//time.Sleep(time.Second)
				fmt.Println("o")
				//time.Sleep(time.Second)
				//ch <- 666 //写数据
			}

		}()
		//go func() {
		//	for {
		//		<-ch //读数据
		//		fmt.Print("w")
		//		time.Sleep(time.Second)
		//		fmt.Print("o")
		//		time.Sleep(time.Second)
		//		fmt.Print("r")
		//		time.Sleep(time.Second)
		//		fmt.Print("l")
		//		time.Sleep(time.Second)
		//		fmt.Println("d")
		//		time.Sleep(time.Second)
		//		ch <- 666 //写数据
		//	}
		//}()
		ch <- 666 //有协程，会死锁
		ch <- 666 //有协程，会死锁
		ch <- 666 //有协程，会死锁
		ch <- 666 //有协程，会死锁
		time.Sleep(10 * time.Second)
		close(ch)
		fmt.Println("close")
		time.Sleep(10 * time.Second)

	}
	if false {
		fmt.Println("--7.遍历通道---------------")
		var ch = make(chan int, 0)
		//go func() {
		//	ch <- 666 //必须另外开协程写数据，在主协程中写数据，会死锁
		//}()

		go func() {
			for i := 0; i < 5; i++ {
				ch <- i //写数据
			}
			close(ch)
		}()
		for num := range ch {
			fmt.Println("num = ", num)
		}

	}
	if false {
		fmt.Println("--8.timer定时器，内部用的是无缓存通道，可实现延时功能，超时处理---------------")
		timer := time.NewTimer(2 * time.Second)
		fmt.Println(time.Now())
		<-timer.C //无缓存通道，阻塞两秒钟
		fmt.Println(time.Now())
		<-time.After(2 * time.Second) //定时两秒，产生一个事件
		fmt.Println(time.Now())
	}
	if false {
		fmt.Println("--9.timer定时器stop---------------")
		timer := time.NewTimer(3 * time.Second)
		go func() {
			<-timer.C //定时器已经停止，阻塞
			fmt.Println("子协程")
		}()
		timer.Stop() //停止定时器
		for {
			time.Sleep(time.Second)
		}
	}
	if false {
		fmt.Println("--9.timer定时器reset---------------")
		timer := time.NewTimer(3 * time.Second)
		timer.Reset(1 * time.Second)
		<-timer.C
		fmt.Println("时间到")
	}
	if false {
		fmt.Println("--10.ticker定时器，内部用的是无缓存通道---------------")
		ticker := time.NewTicker(1 * time.Second)
		i := 0
		for {
			<-ticker.C
			i++
			fmt.Println("i = ", i)
			if i == 5 {
				ticker.Stop()
			}
		}
	}
	if true {
		fmt.Println("--11.select---------------")
		ch := make(chan int)    //数字通信,循环次数
		quit := make(chan bool) //程序是否结束，结束标志
		go func() {
			for i := 0; i < 8; i++ {
				num := <-ch
				fmt.Println(num)
				time.Sleep(10 * time.Second)
			}
			//可以停止
			quit <- true
		}()
		fibonacci(ch, quit)
	}
}
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
			break
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		case <-time.After(3 * time.Second): //如果上面两个case在3秒内执行，当前case永远不会执行
			fmt.Println("超时")
		}
	}
}
