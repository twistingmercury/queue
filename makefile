.PHONY: test bench example

test:
	go test -v . -coverprofile=coverage.out
	go tool cover -html=coverage.out

bench:
	go test -bench=. -benchmem -benchtime=10000x -run=^#

example:
	go run example/main.go