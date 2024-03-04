gen:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/frank269/protoc-gen-code@latest
	go install github.com/complex64/protoc-gen-gorm@latest
	rm -rf generated && mkdir -p generated
	protoc -I ./protos -I ../ -I ../proto --code_out="paths=source_relative:./generated" --go_out="paths=source_relative:./generated" --go-grpc_out="paths=source_relative:./generated" --gorm_out="paths=source_relative:./generated" protos/*.proto;

clean:
	rm -rf generated && mkdir -p generated

run:
	go run main.go