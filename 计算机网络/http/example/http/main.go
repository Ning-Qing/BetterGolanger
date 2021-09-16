package main

import (
	"log"
	"net/http")


func f(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/",f)
	err := http.ListenAndServe(":8080",nil)
	if err !=nil{
		log.Fatal(err)
	}
}