package main

import (
	"fmt"
)

func reverse(ss string) string {
	s := []rune(ss)
	if len(s) < 2 {
		return ss
	}

	start, end := 0, len(s)-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start += 1
		end -= 1
	}
	return string(s)
}

func main() {
	fmt.Println(reverse("abcd"))
}
