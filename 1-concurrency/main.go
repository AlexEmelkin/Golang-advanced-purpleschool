package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sh := make(chan int)
	wg.Add(1)
	go func() {
		square(sh)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(sh)
	}()

	for s := range sh {
		fmt.Print(s, " ")
	}

}

func ran(ch chan int) {
	defer close(ch)
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	for _, a := range arr {
		ch <- a
	}

}

func square(sh chan int) {

	ch := make(chan int, 10)
	go ran(ch)
	for x := range ch {
		y := x * x
		sh <- y
	}

}
