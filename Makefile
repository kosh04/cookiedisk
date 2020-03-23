cookiedisk:
	go build ./cmd/$@

run:
	go run ./cmd/cookiedisk/main.go $(ARGS)

test:
	go vet ./...
	go test -v ./...

.PHONY: run test
