package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Printf("%v, %T\n", f, f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Printf("%v, %T\n", i, i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Printf("%v, %T\n", d, d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Printf("%v, %T\n", u, u)

	k, _ := strconv.Atoi("135")
	fmt.Printf("%v, %T\n", k, k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
