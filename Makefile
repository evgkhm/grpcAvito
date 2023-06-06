.PHONY: \
	proto \
	run \
	lint \

proto:
	protoc -I ./proto --go_out ./internal/service/spec \
	--go_opt paths=source_relative \
	--go-grpc_out ./internal/service/spec \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./internal/service/spec \
	 --grpc-gateway_opt paths=source_relative ./proto/server.proto

run:
	docker-compose up --build

lint:
	golangci-lint cache clean
	golangci-lint run --config=./.golangci.yaml


