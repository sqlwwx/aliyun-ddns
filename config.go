package main

import (
	"github.com/Unknwon/goconfig"
	"log"
)

const CONFIG_FILE = "config.ini"

type Config struct {
	AccessKeyId  string `AccessKeyId`
	AccessKeySec string `AccessKeySec`
	DnsServer    string `DnsServer`
	DomainName   string `域名`
	RR           string `RR`
	RecordId     string `RecordId`
	UsePublicIp  bool   `使用公网ip`
}

func loadConfig() Config {
	config = Config{}
	cfg, err := goconfig.LoadConfigFile(CONFIG_FILE)
	if err != nil {
		log.Println("load config failed")
	}
	config.AccessKeyId, _ = cfg.GetValue("aliyun-ddns", "AccessKeyId")
	config.AccessKeySec, _ = cfg.GetValue("aliyun-ddns", "AccessKeySec")
	config.DnsServer, _ = cfg.GetValue("aliyun-ddns", "DnsServer")
	config.DomainName, _ = cfg.GetValue("record", "DomainName")
	config.RR, _ = cfg.GetValue("record", "RR")
	config.RecordId, _ = cfg.GetValue("record", "Id")
	config.UsePublicIp, err = cfg.Bool("record", "UsePublicIp")
	if err != nil {
		log.Println("load config failed", err)
	}
	return config
}
