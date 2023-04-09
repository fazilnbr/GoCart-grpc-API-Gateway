SHELL := /bin/sh
proto: ## To generate the grpc protocols
	protoc pkg/auth/pb/*.proto --go_out=plugins=grpc:.
	# protoc pkg/auth/pb/*.proto --go_out=. --go-grpc_out=.


run: ## To run the api gateway
	go run cmd/main.go