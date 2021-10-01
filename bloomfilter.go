package main

import (
	"fmt"
)

func h1(word string) int {
	h := 0
	for _, c := range word {
		h += int(c)
	}
	return h
}

func h2(word string) int {
	h := 0
	for _, c := range word {
		n := int(c)
		h += n * n
	}
	return h
}

func h3(word string) int {
	h := 0
	for _, c := range word {
		n := int(c)
		h += n * n * n
	}
	return h
}

func main() {
	var w string
	words := []string{}
	bf := 0
	size := 23
	for {
		fmt.Printf("> ")
		fmt.Scan(&w)

		if w == "exit" {
			break
		}

		h1, h2, h3 := h1(w), h2(w), h3(w)
		mask := 0
		mask |= 1 << (h1 % size)
		mask |= 1 << (h2 % size)
		mask |= 1 << (h3 % size)
		if mask&bf == 0 {
			words = append(words, w)
			bf |= mask
		} else {
			fmt.Printf("%q may already be present\n", w)
		}
		fmt.Printf("%v\n\n", words)
	}
}
