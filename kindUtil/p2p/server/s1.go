package main

import (
	"fmt"
	"net"
)

var peers = make(map[string]net.Addr)

func main() {
	// 创建一个 UDP 监听
	listen, err := net.ListenPacket("udp", ":8081")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	c := make(chan int)
	go sendToEachOther(c)
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
		peers[msg] = addr
		if len(peers) == 2 {
			c <- 1
		}
	}
}

func sendToEachOther(c chan int) {
	for {
		select {
		case <-c:

		}

	}
}
