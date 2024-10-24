test:
	echo 'test'

run:
	go run ./cmd/simpleserver/main.go

build:
	go build -o ./bin/simpleserver ./cmd/simpleserver/main.go
