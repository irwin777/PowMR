build:
	git pull
	go build -o bin/test main.go
install:build
	systemctl stop ipwm3000
	rm test
	cp bin/test .
	systemctl start ipwm3000
	systemctl status ipwm3000