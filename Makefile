build:
	go build -o bin/test main.go
install: build
	scp bin/test ../
