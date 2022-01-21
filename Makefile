.PHONY: proto doc server client .IGNORE cleanServer cleanClient build runs runc

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

appName=proxy
dockerContainerNameServer=proxyServer
dockerContainerNameClient=proxyClient

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0


tag=`git branch | grep \* | cut -d ' ' -f2`

server: cleanServer .IGNORE build runs

build:
	# 先编译
	go build -o ${appName} cmd/main.go
	# 编译镜像
	docker build -t ${appName}:${tag} -f ./Dockerfile .
	docker run -it -d --name ${dockerContainerNameServer} --network host ${appName}:${tag} /cmd/${appName} server

cleanServer:
	@-docker rm -f ${dockerContainerNameServer}

runs:
	docker run --rm -it -d --name ${dockerContainerNameServer} --network host ${appName}:${tag} /cmd/${appName} server