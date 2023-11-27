generate:
	- protoc -I proto proto/sso/sso.proto --go_out=proto/gen/go --go_opt=paths=source_relative --go-grpc_out=proto/gen/go --go-grpc_opt=paths=source_relative
