SHELL := /bin/sh
proto: ## To generate the grpc protocols
	protoc pkg/auth/pb/*.proto --go_out=plugins=grpc:.
	# protoc pkg/auth/pb/*.proto --go_out=. --go-grpc_out=.

path:
	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export GOBIN=$GOPATH/bin
	export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
	export PATH=$PATH:/usr/local/go/bin


swag: ## Generate swagger2 docs
	swag init -g pkg/auth/routes/register.go --parseDependency -o ./cmd/api/docs

run: ## To run the api gateway
	go run cmd/main.go

