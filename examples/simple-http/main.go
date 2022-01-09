package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	os.Setenv("VERSION", "V0.1.0")
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	//获取请求头，设置响应头
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	//设置环境变量到请求头
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	//设置状态码
	// w.WriteHeader(200)
	//设置ResponseBody
	io.WriteString(w, "200")
	//输出IP
	fmt.Printf("request ip: %s; response status code: %d\n", RemoteIp(r), http.StatusOK)
	// to-do 获取状态码
}

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
