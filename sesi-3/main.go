package main

import (
	"fmt"
	"strings"
)

func main() {
	var word = "selamat malam"
	words := strings.Split(word, "")

	maps := make(map[string]int)
	for _, v := range words {
		fmt.Println(v)

		_, ok := maps[v]
		if !ok {
			maps[v] = 1
		} else {
			maps[v]++
		}
	}

	fmt.Println(maps)
}
