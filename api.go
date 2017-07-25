package main

import (
	"github.com/denverdino/aliyungo/dns"
	"log"
)

func List(domain string) {
	client := dns.NewClient(config.AccessKeyId, config.AccessKeySec)
	res, err := client.DescribeDomainRecords(&dns.DescribeDomainRecordsArgs{
		DomainName: domain,
	})
	if err != nil {
	}

	for _, v := range res.DomainRecords.Record {
		log.Println(v.RecordId, v.RR, v.Value, v.DomainName, v.Type)
	}
}

func Update(recordId, rr, value string) error {
	client := dns.NewClient(config.AccessKeyId, config.AccessKeySec)
	_, err := client.UpdateDomainRecord(&dns.UpdateDomainRecordArgs{
		RecordId: recordId,
		RR:       rr,
		Value:    value,
		Type:     dns.ARecord,
	})
	return err
}
