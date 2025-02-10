package helper

import (
	"net"
	"os"
	"runtime"
)

func GetDeviceInfo() string {
	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hostname
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		// Check if the address is an IP address and not a loopback
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil { // IPv4 check
				return ipNet.IP.String()
			}
		}
	}

	return ""
}

func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
