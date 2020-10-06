package znet

import "github.com/kaidev1024/gokai/zinx/ziface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request ziface.IRequest)  {}
func (br *BaseRouter) Handle(request ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
