@echo off
echo Generating protobuf code...

protoc --go_out=gen/go/v1 --go_opt=paths=source_relative --go-grpc_out=gen/go/v1 --go-grpc_opt=paths=source_relative --proto_path=proto proto/payment/service_payment_manager.proto

echo Done!