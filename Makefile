build:
	go build -o bin/main server.go

start:
	go run server.go

dev:
	gin -appPort 3000 --all -i run server.go