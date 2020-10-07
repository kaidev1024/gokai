package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(l uint32) {
	m.DataLen = l
}
