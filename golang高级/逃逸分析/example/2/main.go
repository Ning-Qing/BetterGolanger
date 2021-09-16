package main

func f(n int)int{
	a := make([]int,n)
	l:=len(a)
	return l
}

func main() {
	a := [13107200]int64{}
	a[0]=0
	b := [131072]int64{}
	b[0]=0
	f(100)
}