# test-dev-golang-api

## Structure (hexagonal architecture)

```
├── adapter
│   ├── primary
│   │   ├── api
|   │   │   ├── api.go
|   │   │   ├── route.go
|   │   │   └── swagger.go
│   │   └── handler
|   │   │   ├── client.go
|   │   │   └── handler.go
│   └── secondary
│   │   ├── database
|   │   │   ├── client.go
|   │   │   ├── postgres.go
|   │   │   └── sql.go
│   │   └── rabbitMQ
|   │   │   └── rabbitmq.go
├── cmd
│   └── main.go
├── config
│   ├── docs
|   │   ├── docs.go
|   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── .env
│   ├── .env-example
│   └── config.go
├── deployment
│   ├── cloud-build.yaml
│   └── Dockerfile.go
├── internal
│   ├── entity
│   │   └── entity.go
│   ├── helper
│   │   └── serialize.go
│   ├── service
│   │   ├── client.go
│   │   └── service.go
├── script
│   └── migrations
│   │   ├── 00001_create_table_a.up.sql
│   │   └── 00001_create_table_a.down.sql
├── go.mod
├── go.sum
```

run swagger : http://localhost:YOUR-PORT/v1/docs/index.html