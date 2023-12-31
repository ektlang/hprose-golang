package main

import "github.com/ektlang/hprose-golang/rpc"

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	server := rpc.NewUnixServer("unix:/tmp/my.sock")
	server.AddFunction("hello", hello)
	server.Start()
}
