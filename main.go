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
		} else if command == "get" {
			resp, _ := FindById(config.RecordId)
			log.Println(resp.Value)
		} else if command == "publicIp" {
			log.Println(GetPublicIp())
		} else if command == "internalIp" {
			log.Println(GetLocalIP())
		}
	} else {
		checkAndUpdate()
	}
}

func checkAndUpdate() {
	domainRecord, err := FindById(config.RecordId)
	if err != nil {
		panic(err)
	}
	var currentIp = ""
	if config.UsePublicIp {
		log.Println("use public ip")
		currentIp = GetPublicIp()
	} else {
		log.Println("use internal ip")
		currentIp = GetLocalIP()
	}
	log.Println(domainRecord.Value, currentIp)
	if domainRecord.Value == currentIp {
		log.Println("Need not to update, current IP: ", currentIp)
	} else {
		log.Println("startUpdate")
		log.Println(Update(config.RecordId, config.RR, currentIp))
	}
}
