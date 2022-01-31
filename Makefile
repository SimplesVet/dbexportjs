test-all:
	go test ./...
export-all:
	go run ./cmd/main.go all
observe:
	go run ./cmd/main.go observe
build:
	go build -o dbexport ./cmd
