package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 100000000)
	for i := 0; i < 100000000; i++ {
		wg.Add(1)
		go doSomething(i, &wg, ch)
	}
	wg.Wait()
	fmt.Println("all done")
	for i := 0; i < 100000000; i++ {
		dd := <-ch
		fmt.Println("from ch:"+strconv.Itoa(dd))
	}
}

func doSomething(index int, wg  *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Println("start done:" + strconv.Itoa(index))
	//time.Sleep(20 * time.Millisecond)
	ch <- index
}
