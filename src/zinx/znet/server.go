package znet

import (
	"fmt"
	"net"
	"zin/src/zinx/ziface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (this *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP: %s, Port: %d, is starting\n", this.IP, this.Port)
	go func() {
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(this.IPVersion, fmt.Sprintf("%s:%d", this.IP, this.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		//2 监听服务的地址
		listen, err := net.ListenTCP(this.IPVersion, addr)
		if err != nil {
			fmt.Println("listen: ", this.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server success", this.Name, "Listening...")
		//3 阻塞的等待客服端的连接，处理客户端连接的业务
		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("rec err", err)
						continue
					}

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()
}

func (this *Server) Stop() {
	//TODO 将一些服务器的资源、状态或者一些已经开始的服务结束
}

func (this *Server) Serve() {
	this.Start()
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7523,
	}
	return s
}
