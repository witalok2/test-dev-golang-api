# Export enviroment variables to commands
export

# Variables
go_cover_file=coverage.out

up:: ## Up containers from Docker compose
	@ docker-compose up -d

down:: ## Down local development enviroment
	@ docker-compose down --remove-orphans

test:: ## Do the tests in go
	@ go test -race -coverprofile $(go_cover_file) ./...

swag:
	swag init -g cmd/main.go -o ./config/docs