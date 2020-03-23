cookiedisk:
	go build ./cmd/cookiedisk/ -o $@

run:
	go run ./cmd/cookiedisk/main.go $(ARGS)

test:
	go test -v ./...

.PHONY: run test
