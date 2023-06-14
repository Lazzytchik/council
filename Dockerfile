FROM golang:1.20.4

WORKDIR /app

RUN mkdir -p /app/src
COPY . /app

RUN go get "github.com/go-redis/redis"
# RUN go get "go.mongodb.org/mongo-driver/mongo"
RUN go get "google.golang.org/grpc"
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD make serve

EXPOSE 9000