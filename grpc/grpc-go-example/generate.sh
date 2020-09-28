go get github.com/golang/protobuf/protoc-gen-go
protoc src/calculatorpb/calculator.proto --go_out=plugins=grpc:.
