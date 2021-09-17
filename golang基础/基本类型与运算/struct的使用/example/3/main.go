package main

import (
	"fmt"
)


type A struct {
	name string
}

func (a A) Name() string {
	return a.name
}

func Name(a A) string {
	return a.name
}

func (a *A)SetName(name string){
	a.name = name
}

func main() {
	fmt.Printf("A.Name() type is %T \n",A.Name)
	fmt.Printf("Name(A) type is %T \n",Name)

	a := A{name:"golang"}
	p := &a
	a.SetName("gin")
	fmt.Println(p.Name())

	// A{name:"golang"}.SetName("gin") cannot call pointer method on A{...}
	(&A{name:"golang"}).SetName("gin")

	f1 := A.Name
	f1(a)

	f2 := a.Name
	f2()
}