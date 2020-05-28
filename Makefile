GO_CMD = go
VARSION = -u

all:  fmt tools clean build

tools:
	$(GO_CMD) get $(VARSION)

build:
	$(GO_CMD) build -o ./target/ginapi ./cmd/ginapi/main.go

fmt:
	$(GO_CMD) fmt ./...

clean:
	rm -rf ./target