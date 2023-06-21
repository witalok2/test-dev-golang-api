package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
)

type (
	Environment struct {
		API
		Database
		Queue
	}
	API struct {
		Port string
	}
	Database struct {
		URI string
	}
	Queue struct {
		URI          string
		QueueName    string
		RoutingKey   string
		ExchangeName string
		ExchangeType string
	}
)

func LoadConfig() (*Environment, error) {
	logger.Info("Loading environment variables")

	gaeEnv := os.Getenv("GAE_ENV")
	if gaeEnv == "" {
		//Case it's running locally
		err := godotenv.Load("./config/.env")
		if err != nil {
			return nil, errors.New("error load env locally")
		}
	}

	defaultPort, ok := os.LookupEnv("PORT")
	if !ok {
		defaultPort = "8080"
	}

	databaseURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return nil, errors.New("env var isn't set: DATABASE_URL")
	}

	rabbitURI, ok := os.LookupEnv("RABBITMQ_URI")
	if !ok {
		return nil, errors.New("env var isn't set: RABBITMQ_URI")
	}

	queueName, ok := os.LookupEnv("QUEUE_NAME")
	if !ok {
		return nil, errors.New("env var isn't set: QUEUE_NAME")
	}

	routingKey, ok := os.LookupEnv("ROUTING_KEY")
	if !ok {
		return nil, errors.New("env var isn't set: ROUTING_KEY")
	}

	exchangeName, ok := os.LookupEnv("EXCHANGE_NAME")
	if !ok {
		return nil, errors.New("env var isn't set: EXCHANGE_NAME")
	}

	exchangeType, ok := os.LookupEnv("EXCHANGE_TYPE")
	if !ok {
		return nil, errors.New("env var isn't set: EXCHANGE_TYPE")
	}

	logger.Info("Successfully loaded all environment variables")

	return &Environment{
		API: API{
			Port: defaultPort,
		},
		Database: Database{
			URI: databaseURL,
		},
		Queue: Queue{
			URI:          rabbitURI,
			QueueName:    queueName,
			RoutingKey:   routingKey,
			ExchangeName: exchangeName,
			ExchangeType: exchangeType,
		},
	}, nil
}
