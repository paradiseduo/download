package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

var path = flag.String("path", ".", "创建文件服务器的路径")
var port = flag.Int("port", 23456, "端口号")

func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "0.0.0.0"
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "0.0.0.0"
}

func main() {
	flag.Parse()
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*path))
	loggingFileServerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		if ipPort := strings.Split(clientIP, ":"); len(ipPort) > 1 {
			clientIP = ipPort[0]
		}
		log.Printf("Accessed path: %v from IP: %v", r.URL.Path, clientIP)
		fileServer.ServeHTTP(w, r)
	})

	mux.Handle("/", loggingFileServerHandler)

	log.Printf("Server starting at %s:%d", getIPAddress(), *port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), mux)
	if err != nil {
		log.Fatal(err)
	}
}
