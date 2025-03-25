FROM golang:1.24.1 as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /mktextr ./cmd/mktextr

FROM alpine:latest as run

RUN apk --no-cache add ca-certificates

COPY --from=build /mktextr /app/mktextr

WORKDIR /app

RUN chmod +x /app/mktextr

CMD ["/bin/sh", "-c", "/app/mktextr --domain $DOMAIN"]
