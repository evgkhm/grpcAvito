FROM golang:1.19.3

RUN go version
ENV GOPATH=/

# Expose our port to the world
EXPOSE 50051
EXPOSE 8888

COPY ./ ./

RUN go mod download
RUN go build -o wallet-app ./cmd/main.go

CMD ["./wallet-app"]