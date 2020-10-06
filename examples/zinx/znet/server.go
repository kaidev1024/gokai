package znet

import (
	"errors"
	"fmt"
	"net"

	"github.com/kaidev1024/gokai/zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    ziface.IRouter
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

			dealConn := NewConnection(conn, cid, s.Router)
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

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("add router success")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:	   nil
	}
	return s
}
