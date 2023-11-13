package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	// "unicode/utf8"
	// "bytes"
)

type Config struct {
	Host     string `toml:"Host"`
	Port     string `toml:"Port"`
	Username string `toml:"Username"`
	Password string `toml:"Password"`
	DBName   string `toml:"DBName"`
	SSLMode  string `toml:"SSLMode"`
}

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

	// sl := make([]int, 2, 3)
	// sl = append(sl, 2)
	// fmt.Fprintf("%p", &sl)
	// sl = append(sl, 4)
	// sl[4] = 3
	// fmt.Println(cap(sl))


	var cfg Config

	toml.DecodeFile("qwer/qwe.toml",&cfg)
	fmt.Println(cfg.Host)
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}
