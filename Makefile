.PHONY: proto doc

proto:
	@truss --lib ./proto/lib,../ ./proto/msg/v1/msg.proto
	@truss --lib ./proto/lib,../ --svcout ./kit/users/v1 ./proto/users/v1/users.proto
	@truss --lib ./proto/lib,../ --svcout ./kit/clients/v1 ./proto/clients/v1/clients.proto
	@truss --lib ./proto/lib,../ --svcout ./kit/transfers/v1 ./proto/transfers/v1/transfers.proto
	@truss --lib ./proto/lib,../ --svcout ./kit/tunnels/v1 ./proto/tunnels/v1/tunnels.proto

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