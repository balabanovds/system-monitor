FROM golang:1.15-alpine as builder
WORKDIR /app
COPY . .
RUN rm -rf client && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

### second stage
FROM alpine:3
WORKDIR /app
COPY --from=builder /app/main .
COPY ./config .
EXPOSE $GRPC_PORT
EXPOSE $HTTP_PORT
CMD [ "./main", "-json", "config.json" ]