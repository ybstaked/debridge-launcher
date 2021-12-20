FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./cmd/orbitdb/main.go

FROM alpine
COPY --from=builder /build/main /app/
COPY --from=builder /build/app/orbitdb/config.yml /app/
COPY --from=builder /build/app/orbitdb/ipfs_config /app/

WORKDIR /app
EXPOSE 3000
EXPOSE 4001
CMD ["./main"]
