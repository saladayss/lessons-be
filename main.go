package main

import "fmt"

func main() {
	var (
		i int     = 1
		f float32 = 1.1
	)

	f = float32(i)
	fmt.Println(i, f)
}
