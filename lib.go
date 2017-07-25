package main

import (
	"log"
	"net"
	"os/exec"
	"strings"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func GetDomainIp(dnsServer string) string {
	cmd := exec.Command(`dig`, `@223.5.5.5`, config.RR+"."+config.DomainName, `+short`)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.TrimSpace(string(out))
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "223.5.5.5:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}
