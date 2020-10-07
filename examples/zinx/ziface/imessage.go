package ziface

type IMessage interface {
	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte
	SetMsgId(uint32)
	SetMsgData([]byte)
	SetDataLen(uint32)
}
