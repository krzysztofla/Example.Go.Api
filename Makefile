swagger: 
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
build:
	go build -o ./build ./cmd/server