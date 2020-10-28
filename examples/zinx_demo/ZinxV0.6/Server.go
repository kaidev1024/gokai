package main

import (
	"fmt"

	"github.com/kaidev1024/gokai/zinx/ziface"
	"github.com/kaidev1024/gokai/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router ping handle")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping...\n"))
	// if err != nil {
	// 	fmt.Println("callback  ping error")
	// }
	fmt.Println("recv from client: msgID = ", request.GetMsgID(),
		", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(0, []byte("ping...ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (this *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router hello handle")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping...\n"))
	// if err != nil {
	// 	fmt.Println("callback  ping error")
	// }
	fmt.Println("recv from client: msgID = ", request.GetMsgID(),
		", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("hello..."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[zinx V0.6]")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
