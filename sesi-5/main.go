package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go firstProcess(i, &wg)
		go secondProcess(i, &wg)
	}

	wg.Wait()
}

func firstProcess(index int, wg *sync.WaitGroup) {
	data := []string{"bisa1", "bisa2", "bisa3"}
	fmt.Println(data, index)

	wg.Done()
}

func secondProcess(index int, wg *sync.WaitGroup) {
	data2 := []string{"coba1", "coba2", "coba3"}
	fmt.Println(data2, index)

	wg.Done()
}
