package main

import "fmt"

func isValid(s string) bool {
	m :=
	b := []byte(s)
	stack := make([]byte,0,len(b))
	for len(b)>0{
		char := b[len(b)-1]
		if len(stack)==0 || stack[len(stack)-1]!=char+1{
			stack=append(stack,char)
			b = b[:len(b)-1]
		}else{
			stack=stack[:len(stack)-1]
			b = b[:len(b)-1]
		}
	}
	if len(stack)==0&&len(b)==0{
		return true
	}
	return false
}

func main() {
	
	s :="(){}[]"
	fmt.Println(isValid(s))

}