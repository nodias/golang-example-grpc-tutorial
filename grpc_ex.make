GRPCDIR=grpc/v1.0
CLIENT=cmd/client
SERVER=cmd/server
OUTPUT=out

all : clean build

clean:
	rm -rf $(OUTPUT)

build:
	go build -o $(OUTPUT) $(SERVER)/main.go

call:
	go build -o $(OUTPUT) $(CLIENT)/main.go

proto:
	protoc --go_out=plugins=grpc:. $(GRPCDIR)/helloworld.proto