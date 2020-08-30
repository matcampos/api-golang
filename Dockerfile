
FROM golang:1.14

# Set the Current Working Directory inside the container
WORKDIR /go/src/api-golang

COPY . .

RUN go get -d -v ./...

RUN go run main.go

EXPOSE 3000

CMD ["api-golang"]