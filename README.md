# download
golang实现的一个简单文件服务器，可以快速搭建一个内网文件传输服务

```bash
❯ go build -o download main.go
❯ chmod +x download
❯ cp download /usr/local/bin
❯ download -path . -port 34567
2024/12/24 15:51:49 Server starting at 192.168.123.123:34567
```

在浏览器中打开 http://192.168.123.123:34567/
