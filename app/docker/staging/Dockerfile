FROM golang:1.17-alpine AS gobuilder
WORKDIR /go/src/github.com/eldhoral/eldho-kuncie
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /paylater-cms-api ./main.go

FROM alpine:latest
RUN apk update && apk add ca-certificates bash jq curl tzdata && rm -rf /var/cache/apk/*
ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=gobuilder /paylater-cms-api /bin
COPY --from=gobuilder /go/src/github.com/eldhoral/eldho-kuncie/ /go/src/github.com/eldhoral/eldho-kuncie/
ENTRYPOINT ["/bin/paylater-cms-api", "http"]

