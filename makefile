build:
	go build -o ./bin/panda-blockchain
run: build
	./bin/panda-blockchain
test:
	go test -v ./...