run:
	go run ./cmd/main.go

test:
	go test

build:
	go build -o gopamin ./cmd/main.go

install: 
	go install ./cmd/main.go