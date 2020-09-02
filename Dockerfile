
FROM golang:1.14

# Set the Current Working Directory inside the container
WORKDIR /go/src/api-golang

ADD . /go/src/api-golang

RUN cp docker.env .env

RUN go get -d -v ./...

RUN go install api-golang

EXPOSE 3000

CMD ["api-golang"]