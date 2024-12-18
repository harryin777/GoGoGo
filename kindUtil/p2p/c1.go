package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	//设定参数
	if len(os.Args) < 5 {
		fmt.Println("./client tag remoteIP remotePort")
		return
	}
	//本地绑定端口
	port, _ := strconv.Atoi(os.Args[4])
	//客户端标识
	tag := os.Args[1]
	//服务器IP
	remoteIP := os.Args[2]
	//服务器端口
	remotePort, _ := strconv.Atoi(os.Args[3])
	//绑定本地端口
	localAddr := net.UDPAddr{Port: port}
	//与服务器建立联系
	conn, err := net.DialUDP("udp", &localAddr, &net.UDPAddr{IP: net.ParseIP(remoteIP), Port: remotePort})
	if err != nil {
		log.Panic("UDP拨号失败")
	}

	//发送消息，提供身份
	conn.Write([]byte("I am SEVERA" + tag))
	//从服务器中获得目标地址
	buf := make([]byte, 256)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Panic("读取消息失败", err)
	}
	conn.Close()
	toAddr := parseAddr(string(buf[:n]))
	p2p(&localAddr, &toAddr)
}

func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}

func p2p(srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) {
	//请求建立联系
	conn, _ := net.DialUDP("udp", srcAddr, dstAddr)
	//发送打洞消息
	conn.Write([]byte("打洞消息"))
	//启动goroutine监控标准输入
	go func() {
		buf := make([]byte, 256)
		for {
			//接受UDP消息打印
			n, _, _ := conn.ReadFromUDP(buf)
			if n > 0 {
				fmt.Printf("收到消息:%sp2p>", buf[:n])
			}

		}
	}()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("p2p>")
		//读取标准输入，以换行为读取标志
		data, _ := reader.ReadString('\n')
		conn.Write([]byte(data))
	}
}
