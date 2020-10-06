go get \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go

protoc \
-I. \
-I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.0/third_party/googleapis \
--go_out=plugins=grpc:. \
--grpc-gateway_out=logtostderr=true:. \
--swagger_out . \
src/calculatorpb/calculator.proto
