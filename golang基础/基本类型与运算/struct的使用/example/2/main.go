package main

import (
	"fmt"
	"unsafe"
)

type t1 struct {
	int8
	int64
	int16
	int32
	bool
}

type t2 struct {
	int8
	bool
	int16
	int32
	int64
}

func main() {
	fmt.Printf("int8 size is %v \n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int16 size is %v \n", unsafe.Sizeof(int16(0)))
	fmt.Printf("int32 size is %v \n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int64 size is %v \n", unsafe.Sizeof(int64(0)))

	fmt.Printf("float32 size is %v \n", unsafe.Sizeof(float32(0)))
	fmt.Printf("float64 size is %v \n", unsafe.Sizeof(float64(0)))

	fmt.Printf("bool size is %v \n", unsafe.Sizeof(bool(true)))

	fmt.Printf("string size is %v \n", unsafe.Sizeof("string"))
	fmt.Printf("string size is %v \n", unsafe.Sizeof("string,string"))

	fmt.Printf("slice size is %v \n", unsafe.Sizeof(make([]int, 100, 1000)))
	fmt.Printf("slice size is %v \n", unsafe.Sizeof(make([]string, 100, 1000)))

	fmt.Printf("map size is %v \n", unsafe.Sizeof(make(map[string]int)))
	fmt.Printf("map size is %v \n", unsafe.Sizeof(make(map[string]string)))

	s := "string"
	p := &s
	fmt.Printf("pointer size is %v \n", unsafe.Sizeof(p))

	fmt.Printf("t1 size is %v \n", unsafe.Sizeof(t1{}))
	fmt.Printf("t2 size is %v \n", unsafe.Sizeof(t2{}))
}
