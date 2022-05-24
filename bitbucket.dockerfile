FROM asia-southeast2-docker.pkg.dev/nobu-digital/base-image/golang:1.17-alpine AS builder
ARG nobu_bitbucket_repo_access=$1
ARG LIBRARY_DIR=https://${nobu_bitbucket_repo_access}@bitbucket.org/bitbucketnobubank/go-library.git

RUN apk update && apk add tzdata && apk add git

WORKDIR /app

RUN git clone "$LIBRARY_DIR" "../go-library"

COPY . .

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o paylater-cms-api

###############################
FROM asia-southeast2-docker.pkg.dev/nobu-digital/base-image/golang:1.16-alpine

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/paylater-cms-api /app/paylater-cms-api

RUN mkdir -p /app/logs \
    && chown -R $(id -u $(whoami)):0 /app/logs \
    && chmod -R g+w /app/logs
RUN mkdir -p /app/assets/upload/image \
    && chown -R $(id -u $(whoami)):0 /app/assets/upload/image \
    && chmod -R g+w /app/assets/upload/image

WORKDIR /app
RUN mkdir params
COPY ./migrations/sql/ migrations/sql

# this port should be the same with on env
# EXPOSE 40001

ENTRYPOINT ["/app/paylater-cms-api"]
