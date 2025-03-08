package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// CORS 中间件，用于处理跨域请求
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置允许跨域请求的域名，这里使用 * 表示允许所有域名
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// 设置允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// 处理预检请求（OPTIONS 请求）
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 调用下一个处理函数
		next.ServeHTTP(w, r)
	})
}

// 处理文件上传的函数
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只支持 POST 请求", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "文件 %s 上传成功", handler.Filename)
}

func main() {
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", 0755)
		if err != nil {
			log.Fatalf("创建 uploads 目录失败: %v", err)
		}
	}

	// 创建一个新的多路复用器
	mux := http.NewServeMux()
	// 注册文件上传处理函数到 /upload 路径
	mux.HandleFunc("/upload", uploadHandler)

	// 提供静态文件访问，访问 `/static/` 时，将请求映射到 `./static/` 目录
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// 使用 CORS 中间件包装多路复用器
	handler := corsMiddleware(mux)

	log.Println("服务器启动，监听端口 7788")
	log.Fatal(http.ListenAndServe(":7788", handler))
}
