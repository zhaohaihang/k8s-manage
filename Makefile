run:
	go run cmd/main.go

build:
	go build cmd/main.go

swag: 
	swag init --pd -d ./cmd,docs

