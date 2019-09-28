GRPCDIR=grpc/v1
CLIENT=cmd/client
SERVER=cmd/server
OUTPUT=out

all : clean build

clean:
	rm -rf $(GRPCDIR)

build:
	go build -o $(OUTPUT) $(SERVER)/main.go

call:
	go build -o $(OUTPUT) $(CLIENT)/main.go

proto:
	protoc --go_out=plugins=grpc:. pb/hello.proto