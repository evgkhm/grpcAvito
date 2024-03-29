.PHONY: \
	proto \
	run \
	lint \

proto:
	protoc -I ./proto --go_out ./internal/service/wallet/spec \
	--go_opt paths=source_relative \
	--go-grpc_out ./internal/service/wallet/spec \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./internal/service/wallet/spec \
	 --grpc-gateway_opt paths=source_relative ./proto/wallet_app.proto ./proto/wallet_app_service.proto

run:
	docker-compose up --build

lint:
	golangci-lint cache clean
	golangci-lint run --config=./.golangci.yaml


