package main

import (
	"fmt"
	"os"
)

/*
	输出其命令行工具
*/
func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		// for i := 1; i < len("abcdefg"); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
