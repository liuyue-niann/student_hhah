package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 设置文件服务目录
	fileDir := "./bin"
	port := ":8080"

	// 确保文件目录存在
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		log.Fatalf("Directory %s does not exist", fileDir)
	}

	// 注册文件下载路由
	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		// 获取请求的文件名
		fileName := filepath.Base(r.URL.Path)
		filePath := filepath.Join(fileDir, fileName)

		// 检查文件是否存在
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		// 打开文件
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "Unable to open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// 设置响应头
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		w.Header().Set("Content-Type", "application/octet-stream")

		// 将文件内容写入响应
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
			return
		}
	})

	// 启动服务器
	fmt.Printf("Starting file download server on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
