.PHONY: run
run:
	go run ./cmd/gopamin/gopamin.go $(filter-out $@,$(MAKECMDGOALS))

test:
	go test

build:
	go build -o gopamin ./cmd/main.go

install: 
	go install ./cmd/gopamin/gopamin.go

%:
	@: