build:
	go build cmd/main.go 

install:
	go install cmd/main.go 

run:
	go run cmd/main.go

install_and_run: install run

build_and_run: build run
