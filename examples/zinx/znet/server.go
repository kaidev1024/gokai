package znet

import (
	"errors"
	"fmt"
	"net"

	"github.com/kaidev1024/gokai/zinx/utils"
	"github.com/kaidev1024/gokai/zinx/ziface"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       int
	MsgHandler ziface.IMsgHandle
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("Write back buf err", err)
		return errors.New("CallBackToClient error")
	}

	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, listener at ID: %s, Port: %d is starting",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version %s, MaxConn: %d, MaxPackageSize: %d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	fmt.Printf("[start] Server Listener at IP: %s, Port: %d, is starting\n", s.IP, s.Port)
	// 1. get a tcp addr
	// 2. listen
	// 3. accept client
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen addr err: ", err)
			return
		}

		fmt.Println("start Zinx server succ", s.Name)
		var cid uint32

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err: ", err)
				continue
			}
			fmt.Println("conn created")

			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++

			go dealConn.Start()
		}
	}()

	select {}
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("add router success")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
	}
	return s
}
