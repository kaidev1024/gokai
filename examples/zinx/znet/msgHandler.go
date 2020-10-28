package znet

import (
	"fmt"
	"strconv"

	"github.com/kaidev1024/gokai/zinx/ziface"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), " is not found! need register!")
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		panic("repeat api, msgID = " + strconv.Itoa(int(msgID)))
	}
	mh.Apis[msgID] = router
	fmt.Println("add api MsgID = ", msgID, " success!")
}
