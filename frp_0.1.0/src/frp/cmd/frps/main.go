package main

import (
	"container/list"
	"io"
	"net"
)

func main() {
	//解析地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:1888")
	// 关键是建立一个tcp监听
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	// 建立一个list用来存放链接
	connList := list.New()
	for {
		// 接受连接
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		// 把连接放入list
		connList.PushBack(conn)
		// 这里为了简化代码就先假定就只有两个连接
		if connList.Len() == 2 {
			// 取出连接,
			clientConn := connList.Back().Value.(*net.TCPConn)
			targetConn := connList.Front().Value.(*net.TCPConn)

			// 进行连接相互转发
			go io.Copy(targetConn, clientConn)
			go io.Copy(clientConn, targetConn)
		}

	}
}
