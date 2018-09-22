package main

import (
	"github.com/denverdino/aliyungo/dns"
	"log"
)

func List(domain string) {
	client := dns.NewClient(config.AccessKeyId, config.AccessKeySec)
	describeArgs := dns.DescribeDomainRecordsArgs{
		DomainName: domain,
	}
	describeArgs.PageSize = 500
	res, err := client.DescribeDomainRecords(&describeArgs)
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

func FindById(recordId string) (response *dns.DescribeDomainRecordInfoResponse, err error) {
	client := dns.NewClient(config.AccessKeyId, config.AccessKeySec)
	describeArgs := dns.DescribeDomainRecordInfoArgs{
		RecordId: recordId,
	}
	return client.DescribeDomainRecordInfo(
		&describeArgs,
	)
}
