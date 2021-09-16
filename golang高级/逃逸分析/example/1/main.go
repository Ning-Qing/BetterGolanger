package main

type A struct {
 Name string
}

func NewAP(name string) *A {
 return &A{Name: name}
}

func NewA(name string) A {
 return A{Name: name}
}

func main() {
 NewAP("golang")
 NewA("golang")
}