.PHONY: licenses
licenses:
	rm -rf ./LICENSES
	go-licenses save ./... --save_path=./LICENSES

.PHONY: start
start:
	rm -f ./cmd/main
	go build -o cmd/main cmd/main.go
	./cmd/main