build:
	CGO_ENABLED=1 CC=arm-linux-gnu-gcc GOOS=linux GOARCH=arm GOARM=7 go build -o bin/test main.go
install: build
	scp bin/test root@192.168.88.225:~/
