package main

import (
	"io"
	"net"
	"time"
)

func main() {

	// 连接到服务器
	serverAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:1888")
	serConn, _ := net.DialTCP("tcp", nil, serverAddr)

	// 连接到本地22端口
	localAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:22")
	localConn, _ := net.DialTCP("tcp", nil, localAddr)

	go io.Copy(serConn, localConn)
	go io.Copy(localConn, serConn)

	// 为了简化代码,这里直接使用等待的方式,来等待连接之间通信,阻止主线程退出.
	time.Sleep(20 * time.Second)

}
