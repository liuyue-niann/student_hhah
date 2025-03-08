package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 1. 设置管理员用户密码
	username := "Administrator" // 或当前用户名
	password := "admin"
	fmt.Println("正在设置管理员密码...")
	setPassword(username, password)

	// 2. 启用远程桌面功能
	fmt.Println("正在启用远程桌面...")
	enableRemoteDesktop()

	// 3. 开启远程访问端口（3389）
	fmt.Println("正在开启远程访问端口 3389...")
	openPort(3389)

	// 4. 检查远程桌面状态
	fmt.Println("正在检查远程桌面状态...")
	checkRemoteDesktopStatus()

	fmt.Println("操作完成！")
}

// 设置管理员用户密码
func setPassword(username, password string) {
	cmd := exec.Command("net", "user", username, password)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("设置密码失败:", err)
		fmt.Println("输出:", string(output))
		return
	}
	fmt.Println("密码设置成功！")
}

// 启用远程桌面功能
func enableRemoteDesktop() {
	cmd := exec.Command("reg", "add", "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Terminal Server",
		"/v", "fDenyTSConnections", "/t", "REG_DWORD", "/d", "0", "/f")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("启用远程桌面失败:", err)
		fmt.Println("输出:", string(output))
		return
	}
	fmt.Println("远程桌面已启用！")
}

// 开启指定端口
func openPort(port int) {
	ruleName := fmt.Sprintf("RemoteDesktop_Port_%d", port)
	cmd := exec.Command("netsh", "advfirewall", "firewall", "add", "rule",
		"name="+ruleName,
		"dir=in",
		"action=allow",
		"protocol=TCP",
		"localport="+fmt.Sprintf("%d", port),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("开启端口失败:", err)
		fmt.Println("输出:", string(output))
		return
	}
	fmt.Printf("端口 %d 已开启！\n", port)
}

// 检查远程桌面状态
func checkRemoteDesktopStatus() {
	cmd := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Terminal Server", "/v", "fDenyTSConnections")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("检查远程桌面状态失败:", err)
		return
	}
	fmt.Println("远程桌面状态:\n", string(output))
}
