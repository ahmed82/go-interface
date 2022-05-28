package main

import (
	"fmt"
	"time"
)

func channel() {
	out1 := make(chan string)
	go task("channel order", out1)
	for msg := range out1 {
		fmt.Println(msg)
	}
	//Or
	/*for {
		msg, open := <-out1
		if !open {
			break
		}
		fmt.Println(msg)
	}*/

}

func task(x string, out chan string) {
	// to make sure you properly closed the channel  use deffer
	/* The defer keyword , when the function it done within the scoop it will excute that code*/
	defer close(out)
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		out <- x
	}
	// close(out) /* moved from manule close to defer*/
	//never close the channel from the reciver
}

func channel2() {
	out1 := make(chan string)
	out2 := make(chan string)
	go func() {
		for {
			time.Sleep(time.Second / 2)
			out1 <- "order Process..."
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			out2 <- "return..."
		}
	}()

	for {
		select {
		case msg := <-out1:
			fmt.Println(msg)
		case msg := <-out2:
			fmt.Println(msg)
		}
	}
	/* the below block will cause a bottleneck// solution is "select "
	for {
		fmt.Println(<-out1)
		fmt.Println(<-out2)
	}*/

}
