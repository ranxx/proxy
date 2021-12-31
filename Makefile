.PHONY: proto doc

proto:
	@protoc -I./proto/msg/v1 -I./proto/lib -I../ -I${GOPATH} --gogofaster_out=paths=source_relative,plugins=grpc:./proto/msg/v1 ./proto/msg/v1/msg.proto
	@protoc -I./proto/users/v1 -I./proto/lib -I../ -I${GOPATH} --gogofaster_out=paths=source_relative,plugins=grpc:./proto/users/v1 ./proto/users/v1/users.proto
	@protoc -I./proto/clients/v1 -I./proto/lib -I../ -I${GOPATH} --gogofaster_out=paths=source_relative,plugins=grpc:./proto/clients/v1 ./proto/clients/v1/clients.proto
	@protoc -I./proto/transfers/v1 -I./proto/lib -I../ -I${GOPATH} --gogofaster_out=paths=source_relative,plugins=grpc:./proto/transfers/v1 ./proto/transfers/v1/transfers.proto
	@protoc -I./proto/tunnels/v1 -I./proto/lib -I../ -I${GOPATH} --gogofaster_out=paths=source_relative,plugins=grpc:./proto/tunnels/v1 ./proto/tunnels/v1/tunnels.proto


doc:
	@protoc --proto_path=. \
           --proto_path=../ \
		   --proto_path=./proto/lib \
		   --proto_path=${GOPATH}/src \
           --openapiv2_out=./docs/swagger \
           --openapiv2_opt=json_names_for_fields=false \
           ./proto/users/v1/users.proto
	@protoc --proto_path=. \
           --proto_path=../ \
		   --proto_path=./proto/lib \
		   --proto_path=${GOPATH}/src \
           --openapiv2_out=./docs/swagger \
           --openapiv2_opt=json_names_for_fields=false \
           ./proto/clients/v1/clients.proto
	@protoc --proto_path=. \
           --proto_path=../ \
		   --proto_path=./proto/lib \
		   --proto_path=${GOPATH}/src \
           --openapiv2_out=./docs/swagger \
           --openapiv2_opt=json_names_for_fields=false \
           ./proto/transfers/v1/transfers.proto
	@protoc --proto_path=. \
           --proto_path=../ \
		   --proto_path=./proto/lib \
		   --proto_path=${GOPATH}/src \
           --openapiv2_out=./docs/swagger \
           --openapiv2_opt=json_names_for_fields=false \
           ./proto/tunnels/v1/tunnels.proto