package netx

import (
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) string {
	// 获取顺序 X-Real-IP => X-Forwarded-For => RemoteAddr
	// 获取 X-Real-IP 或 X-Forwarded-For 头部字段的值
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip
	}

	ip = r.Header.Get("X-Forwarded-For")
	if net.ParseIP(ip) != nil {
		return ip
	}

	// 如果头部字段为空，则从 RemoteAddr 中获取
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	// 如果存在多个 IP 地址，取第一个
	if strings.Contains(ip, ",") {
		ips := strings.Split(ip, ",")
		ip = strings.TrimSpace(ips[0])
	}

	if net.ParseIP(ip) == nil {
		return ""
	}

	return ip
}
