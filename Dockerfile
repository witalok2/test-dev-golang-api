FROM golang:1.17 as build

ENV PORT 8080

WORKDIR /go/src/app

COPY . /go/src/app

RUN go build -o ./cmd/api ./cmd/

ARG DATABASE_URL
ARG RABBITMQ_URI
ARG QUEUE_NAME
ARG ROUTING_KEY
ARG EXCHANGE_NAME
ARG EXCHANGE_TYPE

ENV GAE_ENV=production
ENV DATABASE_URL=$DATABASE_URL
ENV RABBITMQ_URI=$RABBITMQ_URI
ENV QUEUE_NAME=$QUEUE_NAME
ENV ROUTING_KEY=$ROUTING_KEY
ENV EXCHANGE_NAME=$EXCHANGE_NAME
ENV EXCHANGE_TYPE=$EXCHANGE_TYPE

CMD ["/go/src/app/cmd/api"]
