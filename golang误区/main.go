package main

import (
	"fmt"
)

func getSum(n int)int{
	res :=0
	for n>0{
		res +=(n%10)*(n%10)
		n/=10
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _,v := range nums1{
		m[v]=1
	}
	for _,v := range nums2{
		if _,ok:=m[v];ok{
			m[v]++
		}
	}
	res := make([]int,0)
	for k,v:=range m{
		if v>1{
			res=append(res,k)
		}
	}
	return res
}

func main() {
	a :=[...]int{1,2,2,1}
	b :=[...]int{2,2}
	fmt.Println(intersection(a[:],b[:]))

	m :=make(map[int]int)
	m = map[int]int{1:1,2:2}
	fmt.Println(m)

}
