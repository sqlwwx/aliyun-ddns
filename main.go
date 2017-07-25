package main

import (
	"log"
	"os"
)

var config Config

func main() {
	args := os.Args
	config = loadConfig()
	if len(args) > 1 {
		command := args[1]
		if command == "list" {
			List(config.DomainName)
		}
	} else {
		checkAndUpdate()
	}
}

func checkAndUpdate() {
	localIP := GetLocalIP()
	domainIP := GetDomainIp(config.DnsServer)
	log.Println(localIP, domainIP, localIP == domainIP)
	if localIP == domainIP {
		log.Println("Need not to update, current IP: ", localIP)
	} else {
		log.Println("startUpdate")
		log.Println(Update(config.RecordId, config.RR, localIP))
	}
}
