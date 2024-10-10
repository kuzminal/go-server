FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./server ./cmd/server

FROM scratch
COPY --from=builder /app/server /usr/bin/server
ENTRYPOINT [ "/usr/bin/server" ]