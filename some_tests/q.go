package main

import (
	"fmt"
	// "unicode/utf8"
	// "bytes"
)

func main() {
	// nums := make([]int, 1, 2)
	// fmt.Println(nums) // <- what's the output?  0

	// appendSlice(nums, 1024)
	// fmt.Println(nums) // <- what's the output?  0

	// mutateSlice(nums, 1, 512)
	// fmt.Println(nums) // <- what's the output?  1024 512

	// a:="qweйцу"
	// fmt.Println(len(a))
	// fmt.Println(len([]rune(a)))
	// fmt.Println(utf8.RuneCountInString(a))
	// fmt.Println()

	sl := make([]int, 2, 3)
	sl = append(sl, 2)
	fmt.Fprintf("%p", &sl)
	sl = append(sl, 4)
	sl[4] = 3
	fmt.Println(cap(sl))
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
