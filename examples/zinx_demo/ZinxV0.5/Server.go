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
	fmt.Println("call router handle")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping...\n"))
	// if err != nil {
	// 	fmt.Println("callback  ping error")
	// }
	fmt.Println("recv from client: msgID = ", request.GetMsgID(),
		", data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("ping...ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[zinx V0.5]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
