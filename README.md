# Eldho Kuncie

Eldho Kuncie Service

## Prerequisites

**Install Go v 1.17**

Please check the [Official Golang Documentation](https://golang.org/doc/install) for installation.

**Install Mockery**

```bash
go get github.com/vektra/mockery/v2/.../
```

**Upgrade package**

```
- Upgrade single package
go get -u github.com/gorilla/mux
go mod tidy

- Upgrade all
go get -u
go mod tidy

```


## Installation


**Download dependencies (optional)**

If you want to download all dependencies into the vendor folder, please run the following command:

```bash
go mod vendor
```

**Clone this repository**

```bash
git clone github.com/eldhoral/eldho-kuncie.git
# Switch to the repository folder
cd eldho-kuncie
```

**Copy the `.env.example` to `.env`**

```bash
cp .env.example .env
```

Make the required configuration changes in the `.env` file.

## Install mysql on Docker

**Build**

```bash
make docker-build-mysql
```

**Run**

```bash
make docker-run-mysql
```

**Run DB Migration**

```bash
make migrate-sql
```

**Run Application**

```bash
make run
```

## Unit Testing

**Mocking The Interface**
```bash
cd internal/{function folder}
# Mock Repository interface
mockery --name=Repository --output=../mocks
# Mock Service interface
mockery --name=Service --output=../mocks
```

**Run Unit Test**

To run unit testing, just run the command below:
```bash
make test
```

**Code Coverage**

If you want to see code coverage in an HTML presentation (after the test) just run:

```bash
make coverage
```

## Folders

* `cmd` - Contains command files.
* `app/api` - Contains http server.
* `app/docker` - Contains Dockerfile.
* `app/migrations` - Contains DB migrator.
* `internal` - Contains packages which are specific to your project.
* `pkg` - Contains extra packages.

## Reference

* [Folder Explanation](https://github.com/golang-standards/project-layout)
* [Go Modules](https://blog.golang.org/using-go-modules)
* [Google JSON Style Guide](https://google.github.io/styleguide/jsoncstyleguide.xml)
* [Gorilla Mux](https://www.gorillatoolkit.org/pkg/mux)
* [Logrus](https://github.com/sirupsen/logrus)
* [Mockery](https://github.com/vektra/mockery)
* [SQL-Migrate](https://github.com/rubenv/sql-migrate)
* [SQLMock](https://github.com/DATA-DOG/go-sqlmock)
* [Testify](https://github.com/stretchr/testify)

## Contributing

When contributing to this repository, please note we have a code standards, please follow it in all your interactions with the project.

#### Steps to contribute

1. Clone this repository.
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Submit pull request.

**Note** :

* Please make sure to update tests as appropriate.

* It's recommended to run `make test` command before submit a pull request.

* Please update the postman collection if you modify or create new endpoint.
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/16666203-dae0b46c-ef3e-459b-9cf8-e98844fd0e86?action=collection%2Ffork&collection-url=entityId%3D16666203-dae0b46c-ef3e-459b-9cf8-e98844fd0e86%26entityType%3Dcollection%26workspaceId%3D242523f7-daca-4ecd-9e27-d2d288efd49a)

