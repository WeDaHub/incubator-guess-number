install:
	go mod init guess-number

build:
	cd functions/guess-number && GOOS=linux GOARCH=amd64 go build -o main main.go
	