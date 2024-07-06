package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	readAndWriteValue(50)
}

func readAndWriteValue(num int) {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				fmt.Printf("%d - четное\n", i)
			} else {
				fmt.Printf("%d - нечетное\n", i)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start).Seconds())
}
