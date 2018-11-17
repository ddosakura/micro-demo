build:
	protoc -I . --go_out=plugins=grpc:. ./proto/*/*.proto
	# protoc --proto_path=. --micro_out=. --go_out=. ./proto/*/*.proto