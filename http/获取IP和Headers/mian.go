package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func ip(resp http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintln(resp, "IP:", GetClientIP(req))
}

func header(resp http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintln(resp, "req.RemoteAddr:", req.RemoteAddr)
	_, _ = fmt.Fprintln(resp, "req.Host:", req.Host)
	_, _ = fmt.Fprintln(resp, "X-Real-Ip:", req.Header.Get("X-Real-Ip"))
	_, _ = fmt.Fprintln(resp, "X-Forwarded-For:", req.Header.Get("X-Forwarded-For"))
	_, _ = fmt.Fprintln(resp, "X-Forwarded-Proto:", req.Header.Get("X-Forwarded-Proto"))
	_, _ = fmt.Fprintln(resp, "X-Forwarded-Host:", req.Header.Get("X-Forwarded-Host"))
}

// GetClientIP 获取客户端IP
func GetClientIP(req *http.Request) string {
	// 查找 X-Real-Ip
	IPAddress := req.Header.Get("X-Real-Ip")
	// 查找 X-Forwarded-For
	if IPAddress == "" {
		IPS := strings.Split(req.Header.Get("X-Forwarded-For"), ",")
		IPAddress = IPS[0]
	}
	// 查找 RemoteAddr
	if IPAddress == "" {
		IPAddress = req.RemoteAddr
		IPAddress, _, _ = net.SplitHostPort(IPAddress)
	}
	return IPAddress
}

func main() {
	http.HandleFunc("/header", header)
	http.HandleFunc("/ip", ip)

	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
