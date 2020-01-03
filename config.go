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
	Type         string `Type`
	UsePublicIp  bool   `使用公网ip`
}

func loadConfig() Config {
	config = Config{}
	cfg, err := goconfig.LoadConfigFile(CONFIG_FILE)
	if err != nil {
		log.Println("load config failed")
	}
	config.AccessKeyId, _ = cfg.GetValue("aliyun-ddns", "accessKeyId")
	config.AccessKeySec, _ = cfg.GetValue("aliyun-ddns", "accessKeySec")
	config.DnsServer, _ = cfg.GetValue("aliyun-ddns", "dnsServer")
	config.DomainName, _ = cfg.GetValue("record", "domainName")
	config.RR, _ = cfg.GetValue("record", "rr")
	config.RecordId, _ = cfg.GetValue("record", "id")
	config.UsePublicIp = cfg.MustBool("record", "usePublicIp", false)
  config.Type = cfg.MustValue("record", "type", "A")
	return config
}
