package main

import (
	"fmt"
)

func main() {
	// bs_UTF16LE, _, _ := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder(), []byte("測試"))

	// text := []string{"C", "A", "Ш", "A", "P", "B", "O"}
	textMap := make(map[int]string)
	textMap[0] = "C"
	textMap[2] = "A"
	textMap[4] = "Ш"
	textMap[6] = "A"
	textMap[8] = "P"
	textMap[10] = "B"
	textMap[12] = "O"

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	j := 0
	for {
		fmt.Println("Nilai j = ", j)
		j++

		if j == 5 {
			for i, v := range textMap {
				fmt.Printf("character %+q '%s' starts at byte position %d\n", v, v, i)
			}
			continue
		}

		if j == 11 {
			break
		}
	}
}
