package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int64, 10)

	for idx := 0; idx < 10; idx++ {
		go func(idx int) {
			ch <- int64(idx)
		}(idx)
		time.Sleep(500 * time.Millisecond)
	}

	for idx := 0; idx < 10; idx++ {
		data := <-ch
		fmt.Println(data)
	}
}
