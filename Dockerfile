FROM golang:1.19.0-alpine3.16 AS builder

WORKDIR /opt

COPY go.mod go.sum ./
COPY ./vendor ./vendor

RUN go build -v ./vendor/...

COPY ./ ./

RUN go install -v ./cmd/ent
RUN go install -v ./cmd/gorm
RUN go install -v ./cmd/jet
RUN go install -v ./cmd/sqlc

FROM alpine:3.16

COPY --from=builder /go/bin/ent  /bin
COPY --from=builder /go/bin/gorm /bin
COPY --from=builder /go/bin/jet  /bin
COPY --from=builder /go/bin/sqlc /bin

CMD ["/bin/ent"]
