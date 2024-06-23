FROM golang:1.22-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata build-base && update-ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go mod download

RUN go build -o main cmd/api/main.go


FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/main .

ENV TZ=Asia/Jakarta
CMD /main