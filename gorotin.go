package main

import (
	"fmt"
	"sync"
	"time"
)

func goRoutin() {
	var wg sync.WaitGroup
	wg.Add(1)

	/* Anonymous functions */
	go func() {
		/*go*/ tasks("Order")
		//go tasks("refund")
		wg.Done()
	}()
	wg.Wait()
}

func tasks(x string) {
	//for i := 1; true; i++ {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Process", i, x)
	}
}
