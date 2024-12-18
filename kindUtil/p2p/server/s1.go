package main

import (
	"fmt"
	"net"
	"strings"
)

var peers = make([]net.Addr, 0, 10)
var sourceMapTarget = make(map[string]string)

func main() {
	// 创建一个 UDP 监听
	listen, err := net.ListenPacket("udp", ":8081")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	c := make(chan int)
	go sendToEachOther(c, listen)
	// 创建一个缓冲区
	buf := make([]byte, 4096)
	for {
		// 读取数据
		n, addr, err := listen.ReadFrom(buf) // 读取数据包
		if err != nil {
			fmt.Println("read failed, err:", err)
			continue
		}
		msg := string(buf[:n])
		fmt.Printf("收到数据来自 %s: %s\n", addr, msg)
		data := strings.Split(msg, ":")
		if len(data) != 2 {
			listen.WriteTo([]byte("数据有误"), addr)
			continue
		}
		peers = append(peers, addr)
		sourceMapTarget[data[0]] = data[1]
		if len(peers) == 2 {
			c <- 1
		}
	}
}

func sendToEachOther(c chan int, listen net.PacketConn) {
	for {
		select {
		case <-c:
			_, err := listen.WriteTo([]byte(peers[0].String()), peers[1])
			if err != nil {
				panic(err)
			}
			_, err = listen.WriteTo([]byte(peers[1].String()), peers[0])
			if err != nil {
				panic(err)
			}
		}
	}
}
