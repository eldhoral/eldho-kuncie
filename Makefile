run: build
	@./bin/paylater http

test:
	@go fmt ./...
	@go vet ./...
	@go test -v -coverprofile=coverage.out ./...

coverage:
	@go tool cover -html=coverage.out

build:
	@go mod tidy
	@go build -o bin/paylater main.go

engine:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/docker/paylater main.go

image: engine
	@docker build -t paylater .

migrate-sql:
	@$(HOME)/go/bin/sql-migrate up -env="development"

migrate-data:
	@$(HOME)/go/bin/sql-migrate up -env="development"

docker-staging-build:
	@docker-compose -f docker-compose.staging.yml up --build

docker-staging-run:
	@docker-compose -f docker-compose.staging.yml up