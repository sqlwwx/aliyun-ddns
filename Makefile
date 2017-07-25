start:
	go run *.go

build: clean
	go build -ldflags "-s -w" -o ./bin/aliyun_ddns
	upx -9 ./bin/aliyun_ddns

build-darwin: clean
	GOOS=darwin go build -o ./bin/aliyun_ddns-darwin

build-android: clean
	GOOS=linux GOARCH=arm GOARM=7 go build -o ./bin/aliyun_ddns-android

clean:
	rm -f ./bin/*

