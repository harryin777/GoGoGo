package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// 处理连接的函数，既接收又发送消息
func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("qqqqq"))

	reader := bufio.NewReader(conn)

	for {
		// 接收客户端消息
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading from client:", err)
			return
		}
		fmt.Printf("Received from client: %s", message)

		// 向客户端发送响应
		_, err = conn.Write([]byte("Server received: " + message))
		if err != nil {
			log.Println("Error writing to client:", err)
			return
		}
	}
}

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080...")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// 处理连接
		go handleConnection(conn)
	}
}
