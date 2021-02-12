FROM golang:alpine

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main

RUN go get github.com/pilu/fresh
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get -u -v github.com/gobeam/mongo-go-pagination

ENTRYPOINT ["fresh"]