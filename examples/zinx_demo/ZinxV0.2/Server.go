package main

import "github.com/kaidev1024/gokai/zinx/znet"

func main() {
	s := znet.NewServer("[zinx V0.2]")
	s.Serve()
}
