package main

import (
	"fmt"
)

type t1 = int64
type t2 int64

func main() {
	var a t1
	var b t2
	fmt.Printf("t1 type is %T \n", a)
	fmt.Printf("t2 type is %T \n", b)
}
