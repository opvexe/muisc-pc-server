package util

import (
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

/*
	获取远端请求ip
*/
func GetRequestIPAddress(ctx *gin.Context) string {
	ip := ctx.Request.RemoteAddr
	if len(ip) != 0 {
		ip = strings.Split(ip, ":")[0]
	}
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		for _, i := range strings.Split(ctx.Request.Header.Get(h), ",") {
			if len(i) != 0 {
				ip = i
			}
		}
	}
	return ip
}

/*
	获取本机ip
*/
func GetLocalIPAddress() (string, error) {
	ip, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	var ipAdress string
	for _, adr := range ip {
		//检查ip是否回环地址
		if ipnet, ok := adr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { //IPV4
				ipAdress = ipnet.IP.String()
			}
		}
	}
	return ipAdress, nil
}
