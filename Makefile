build:
	GOARCH=arm GOARM=7 go build -o bin/ipwm3000 main.go
install:build
	ssh root@172.20.20.15 "systemctl stop ipwm3000"
	scp bin/ipwm3000 root@172.20.20.15:/root/PowMR/
	ssh root@172.20.20.15 "systemctl start ipwm3000"
	ssh root@172.20.20.15 "systemctl status ipwm3000"