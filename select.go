package main

import "fmt"
import "runtime"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	fmt.Print("hello\t")
	for {
		fmt.Print("send\t")
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Print("world\t")
		for i := 0; i < 10; i++ {
			fmt.Print("before\t")
			fmt.Println(<-c)
			fmt.Print("out\t", i)
		}
		quit <- 0
	}()

	num := runtime.NumCPU()
	fmt.Print(num)
	fibonacci(c, quit)
}
