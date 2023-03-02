FROM golang:1.20.1-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/Johannes-Krabbe/hive-nexus-api/
WORKDIR /go/src/github.com/Johannes-Krabbe/hive-nexus-api
RUN go mod download
COPY . /go/src/github.com/Johannes-Krabbe/hive-nexus-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/hive-nexus-api github.com/Johannes-Krabbe/hive-nexus-api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/Johannes-Krabbe/hive-nexus-api/build/hive-nexus-api /usr/bin/hive-nexus-api
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/hive-nexus-api"]
