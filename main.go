package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	go func(testChan chan int) {
		for i := 0; i < 10; i++ {
			testChan <- i
		}
		close(testChan)
	}(channel)
	for value := range channel {
		fmt.Println(value)
	}
}

func readAndWriteValue(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	buffer := make(chan int)
	go func(chanForWriting chan<- int) {
		for i := 0; i < num; i++ {
			chanForWriting <- i
		}
		close(buffer)
	}(buffer)
	for value := range buffer {
		fmt.Println(value)
	}
}
func readAndWriteValue10(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	buffer := make(chan int)
	go func(chanForWriting chan<- int) {
		for i := 0; i < num; i++ {
			chanForWriting <- i * 10
		}
		close(buffer)
	}(buffer)
	for value := range buffer {
		fmt.Println(value)
	}
}
