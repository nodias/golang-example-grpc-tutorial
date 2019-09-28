GRPCDIR=grpc/v1.0
CLIENT=cmd/client
SERVER=cmd/server
OUTPUT=out

all : clean build

clean:
	rm -rf $(OUTPUT)

build-server:
	env GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT)/server $(SERVER)/main.go

build-client:
	env GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT)/client $(CLIENT)/main.go

proto:
	protoc --go_out=plugins=grpc:. $(GRPCDIR)/helloworld.proto