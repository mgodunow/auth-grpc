generate:
	- protoc -I proto proto/sso/sso.proto --go_out=proto/gen/go --go_opt=paths=source_relative --go-grpc_out=proto/gen/go --go-grpc_opt=paths=source_relative
migrate:
	- go run ./cmd/migrator	--storage-path=../storage/sso.db --migrations-path=../migrations
run:
	- go run ./cmd/sso --config=./config/local.env