package main

import "fmt"

func f1(a [3]int){
	a[0] = 11
	fmt.Println(a)
}

func f2(a *[3]int){
	a[0]=12
	fmt.Println(a)
}

func main() {
	a :=[...]int{1,2,3}
	f1(a)
	f2(&a)
	fmt.Println(a)

	s := a[:]
	
}