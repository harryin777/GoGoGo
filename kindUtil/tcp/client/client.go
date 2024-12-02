package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 启动一个 goroutine 用于接收来自服务器的消息
	go func() {
		reader := bufio.NewReader(conn)
		for {
			// 从服务器接收消息
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading from server:", err)
				return
			}
			fmt.Printf("Received from server: %s", message)
		}
	}()

	// 主线程用于从标准输入读取消息并发送到服务器
	writer := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter message to send: ")
		if scanner.Scan() {
			message := scanner.Text() + "\n"
			_, err := writer.WriteString(message)
			if err != nil {
				log.Println("Error sending message:", err)
				return
			}

			// 确保数据被发送
			writer.Flush()
		}
	}
}
