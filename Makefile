swagger: 
	swagger generate spec -o ./swagger.yaml --scan-models
build:
	go build -o ./build ./cmd/server
run:
	go run ./cmd/server/main.go