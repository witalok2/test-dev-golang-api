package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Simulate the environment variables
	os.Setenv("GAE_ENV", "TEST")
	os.Setenv("PORT", "8080")
	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname")
	os.Setenv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/")
	os.Setenv("QUEUE_NAME", "my-queue")
	os.Setenv("ROUTING_KEY", "my-routing-key")
	os.Setenv("EXCHANGE_NAME", "my-exchange")
	os.Setenv("EXCHANGE_TYPE", "direct")

	defer func() {
		// Clean up the environment variables after the test
		os.Unsetenv("PORT")
		os.Unsetenv("DATABASE_URL")
		os.Unsetenv("RABBITMQ_URI")
		os.Unsetenv("QUEUE_NAME")
		os.Unsetenv("ROUTING_KEY")
		os.Unsetenv("EXCHANGE_NAME")
		os.Unsetenv("EXCHANGE_TYPE")
	}()

	config, err := LoadConfig()
	assert.NoError(t, err)

	expectedConfig := &Environment{
		API: API{
			Port: "8080",
		},
		Database: Database{
			URI: "postgres://user:password@localhost:5432/dbname",
		},
		Queue: Queue{
			URI:          "amqp://guest:guest@localhost:5672/",
			QueueName:    "my-queue",
			RoutingKey:   "my-routing-key",
			ExchangeName: "my-exchange",
			ExchangeType: "direct",
		},
	}

	assert.Equal(t, expectedConfig, config)
}
