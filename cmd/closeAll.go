package main

import (
	"awesomeProject1/core"
	"fmt"
	"strconv"
)

/*
*
全部关机
*/
func main() {
	fmt.Println("正在发送...")
	for i := 1; i < 255; i++ {
		shutdown := core.JiYuSendShutdown("192.168.234."+strconv.Itoa(i), 4705)
		if !shutdown {
			fmt.Println("发送失败:" + "192.168.234." + strconv.Itoa(i))
		}
	}
	fmt.Println("发送成功...")
}
