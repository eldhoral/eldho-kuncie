FROM golang:1.17-alpine AS gobuilder
WORKDIR /go/src/bitbucket.org/bitbucketnobubank/paylater-cms-api


COPY go.mod .
COPY go.sum .
COPY .env .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

COPY . .

ENTRYPOINT CompileDaemon --build="go build -o /go/bin/paylater-cms-api ." --command=/go/bin/paylater-cms-api --directory=.

