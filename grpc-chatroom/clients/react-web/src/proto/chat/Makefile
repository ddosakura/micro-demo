build:
	protoc -I . --go_out=plugins=grpc:. --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. ./*.proto
