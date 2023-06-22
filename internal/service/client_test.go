package service

import (
	"context"
	"sort"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/witalok2/test-dev-golang-api/adapter/secondary/database"
	rabbitmq "github.com/witalok2/test-dev-golang-api/adapter/secondary/rabbitMQ"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
	"github.com/witalok2/test-dev-golang-api/internal/helper"
)

type ServiceTestSuite struct {
	suite.Suite
	ctx context.Context

	databaseMock *database.MockDatabase
	rabbitMQMock *rabbitmq.MockRabbitMQ

	service *service
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.databaseMock = new(database.MockDatabase)
	suite.rabbitMQMock = new(rabbitmq.MockRabbitMQ)
	suite.service = NewService(suite.ctx, suite.databaseMock, suite.rabbitMQMock)
}

func (suite *ServiceTestSuite) TestListClient() {
	cases := map[string]struct {
		err error
	}{
		`list clients`: {
			err: nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.databaseMock.On("ListClient", suite.ctx, 1, 20).Return(&[]entity.Client{}, &entity.Pagination{}, nil)

			clients, pagination, err := suite.service.ListClient(suite.ctx, 1, 20)

			suite.NoError(err)
			suite.NotNil(clients)
			suite.NotNil(pagination)

			suite.databaseMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *ServiceTestSuite) TestGetClient() {
	cases := map[string]struct {
		clientId uuid.UUID
		err      error
	}{
		`get client by id`: {
			clientId: uuid.New(),
			err:      nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			resp := entity.Client{
				ID:   cases[key].clientId,
				Name: "Witalo",
			}

			suite.databaseMock.On("GetClient", suite.ctx, cases[key].clientId).Return(resp, nil)

			client, err := suite.service.GetClient(suite.ctx, cases[key].clientId)

			suite.NoError(err)
			suite.Equal(client.ID.String(), cases[key].clientId.String())
			suite.Equal(client.Name, resp.Name)

			suite.databaseMock.AssertExpectations(suite.T())
		})
	}
}
func (suite *ServiceTestSuite) TestCreateClient() {
	cases := map[string]struct {
		client *entity.Client
		err    error
	}{
		`create client successfully`: {
			client: &entity.Client{
				Name:     "John Doe",
				LastName: "Johnes",
				Contact:  "+5565992526018",
				Address:  "Center",
				Birthday: "1996-02-04",
				CPF:      "05000000000",
			},
			err: nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			message, err := helper.PreperamentQueue(cases[key].client, entity.CREATE_CLIENT)
			suite.NoError(err)

			suite.rabbitMQMock.On("PublishMessage", suite.ctx, message).Return(nil)
			err = suite.service.CreateClient(suite.ctx, cases[key].client)

			suite.NoError(err)
			suite.rabbitMQMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *ServiceTestSuite) TestUpdateClient() {
	cases := map[string]struct {
		client *entity.Client
		err    error
	}{
		`update client successfully`: {
			client: &entity.Client{
				Name:     "Johnes",
				LastName: "John Doe",
				Contact:  "+5565992526018",
				Address:  "Center",
				Birthday: "1996-02-04",
				CPF:      "05000000000",
			},
			err: nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			message, err := helper.PreperamentQueue(cases[key].client, entity.UPDATE_CLIENT)
			suite.NoError(err)

			suite.rabbitMQMock.On("PublishMessage", suite.ctx, message).Return(nil)
			err = suite.service.UpdateClient(suite.ctx, cases[key].client)

			suite.NoError(err)
			suite.rabbitMQMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *ServiceTestSuite) TestDeleteClient() {
	cases := map[string]struct {
		clientId uuid.UUID
		err      error
	}{
		`delete client by id`: {
			clientId: uuid.New(),
			err:      nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.databaseMock.On("DeleteClient", suite.ctx, cases[key].clientId).Return(nil)

			err := suite.service.DeleteClient(suite.ctx, cases[key].clientId)
			suite.NoError(err)

			suite.databaseMock.AssertExpectations(suite.T())
		})
	}
}
