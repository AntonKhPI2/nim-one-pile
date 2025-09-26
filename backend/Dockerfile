FROM golang:1.24.7 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -o server ./cmd/api
FROM alpine:3.20
RUN adduser -D app
WORKDIR /app
RUN apk add --no-cache ca-certificates netcat-openbsd
COPY --from=build /app/server /usr/local/bin/server
COPY docker/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh /usr/local/bin/server
ENV PORT=8080
EXPOSE 8080
USER app
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
