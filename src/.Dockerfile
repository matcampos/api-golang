# FROM golang:alpine
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go build -o main . 
# CMD ["/app/main"]

# FROM golang:1.8.3 as builder
# WORKDIR main
# COPY main.go  .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# CMD ["./main"]

FROM golang:1.8

WORKDIR /go/apiGo/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["main"]