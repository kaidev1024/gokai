package znet

import (
	"fmt"
	"net"

	"github.com/kaidev1024/gokai/zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
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

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err: ", err)
				continue
			}
			fmt.Println("conn created")
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read buf err: ", err)
						continue
					}
					fmt.Printf("received client buf %s\n", buf)
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write buf err: ", err)
						continue
					}
				}
			}()
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

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
