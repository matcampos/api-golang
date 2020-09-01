
FROM golang:1.14

# Set the Current Working Directory inside the container
WORKDIR $GO_PATH/src/matcampos/api-golang

COPY . .

RUN cp .env-docker .env

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main

EXPOSE 3000

RUN chmod -R 777 main

CMD ["./main"]