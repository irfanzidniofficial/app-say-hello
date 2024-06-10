package main

import (
	"fmt"

	go_say_hello "github.com/irfanzidniofficial/go-say-hello/v2"
)

func main() {
	hello:=go_say_hello.SayHello("Irfan")

	fmt.Println(hello)
}