
FROM amd64/golang:1.14

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/matcampos/api-golang/src

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 3000

CMD ["api-bitcoin"]