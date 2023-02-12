run: build
	@./bin/eldho-kuncie http

test:
	@go fmt ./...
	@go vet ./...
	@go test -v -coverprofile=coverage.out ./...

coverage:
	@go tool cover -html=coverage.out

build:
	@go mod tidy
	@go build -o bin/eldho-kuncie main.go

docker-build-mysql:
	@sudo docker pull mysql/mysql-server:latest

docker-run-mysql:
	@docker run --name mysql -e MYSQL_USER=root -e MYSQL_PASSWORD=root -e MYSQL_DATABASE=kuncie -p 3306:3306 -d mysql/mysql-server:latest

migrate-sql:
	@go run . migrate	
