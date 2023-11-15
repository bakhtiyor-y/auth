test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

lint:
	golangci-lint run

BIN:=$(CURDIR)/bin

install:
	GOBIN=$(BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1


gen_auth: ## Генерация proto-файлов
		mkdir -p pkg/auth_v1
		protoc 	--proto_path=api/auth_v1 \
				--proto_path=proto \
				--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
				--plugin=protoc-gen-go=bin/protoc-gen-go \
				--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
				--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
				--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
				--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
				--grpc-gateway_out=pkg/auth_v1 --grpc-gateway_opt=paths=source_relative \
				--validate_out lang=go:pkg/auth_v1 --validate_opt=paths=source_relative \
				--plugin=protoc-gen-validate=bin/protoc-gen-validate \
				--openapiv2_out=allow_merge=true,merge_file_name=api_auth_v1:docs \
                --plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
				api/auth_v1/auth.proto

gen_login: ## Генерация proto-файлов
		mkdir -p pkg/login_v1
		protoc 	--proto_path=api/login_v1 \
				--proto_path=proto \
				--go_out=pkg/login_v1 --go_opt=paths=source_relative \
				--plugin=protoc-gen-go=bin/protoc-gen-go \
				--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
				--go-grpc_out=pkg/login_v1 --go-grpc_opt=paths=source_relative \
				--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
				--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
				--grpc-gateway_out=pkg/login_v1 --grpc-gateway_opt=paths=source_relative \
				--validate_out lang=go:pkg/login_v1 --validate_opt=paths=source_relative \
				--plugin=protoc-gen-validate=bin/protoc-gen-validate \
				api/login_v1/*.proto


