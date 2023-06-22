package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	rabbitmq "github.com/witalok2/test-dev-golang-api/adapter/secondary/rabbitMQ"
	"github.com/witalok2/test-dev-golang-api/config"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
	"github.com/witalok2/test-dev-golang-api/internal/service"
)

type HandlerTestSuite struct {
	suite.Suite
	ctx context.Context

	// database     database.Repository
	serviceMock  *service.MockService
	rabbitMQMock *rabbitmq.MockRabbitMQ
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) SetupSuite() {
	suite.ctx = context.Background()
}

func (suite *HandlerTestSuite) TestListClient() {
	cases := map[string]struct {
		expectedCode int
		expectedList string
		err          error
	}{
		`list clients`: {
			expectedCode: http.StatusOK,
			expectedList: `{"data":[],"metadata":{"page":0,"limit":0,"totalPages":0,"totalItems":0}}`,
			err:          nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.serviceMock = new(service.MockService)
			suite.rabbitMQMock = new(rabbitmq.MockRabbitMQ)

			suite.serviceMock.On("ListClient", suite.ctx, 1, 20).Return(&[]entity.Client{}, &entity.Pagination{}, nil)

			handler := Handler{
				service:     suite.serviceMock,
				environment: &config.Environment{},
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/v1/client?page=1&limit=20", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.ListClient(suite.ctx)(c)
			suite.Require().NoError(err)
			suite.Require().Equal(cases[key].expectedCode, rec.Code)

			suite.JSONEq(cases[key].expectedList, rec.Body.String())
		})
	}
}

func (suite *HandlerTestSuite) TestGetClient() {
	cases := map[string]struct {
		expectedCode int
		expectedBody string
		err          error
	}{
		`get client`: {
			expectedCode: http.StatusOK,
			expectedBody: `{"data":{"id":"647bb6a6-cb08-46f2-a4c7-6d5defa4f573","name":"John Doe","lastName":"Johnes","contact":"+5565992526018","address":"Center","brithday":"1996-02-04","cpf":"05000000000","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}}`,
			err:          nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.serviceMock = new(service.MockService)
			suite.rabbitMQMock = new(rabbitmq.MockRabbitMQ)

			clientID := uuid.MustParse("647bb6a6-cb08-46f2-a4c7-6d5defa4f573")
			client := entity.Client{
				ID:       clientID,
				Name:     "John Doe",
				LastName: "Johnes",
				Contact:  "+5565992526018",
				Address:  "Center",
				Birthday: "1996-02-04",
				CPF:      "05000000000",
			}

			suite.serviceMock.On("GetClient", suite.ctx, clientID).Return(client, nil)

			handler := Handler{
				service:     suite.serviceMock,
				environment: &config.Environment{},
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/v1/client/647bb6a6-cb08-46f2-a4c7-6d5defa4f573", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(clientID.String())

			err := handler.GetClient(suite.ctx)(c)
			suite.Require().NoError(err)
			suite.Require().Equal(cases[key].expectedCode, rec.Code)

			suite.JSONEq(cases[key].expectedBody, rec.Body.String())
			suite.serviceMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *HandlerTestSuite) TestCreateClient() {
	cases := map[string]struct {
		expectedCode int
		requestBody  string
		err          error
	}{
		`create client`: {
			expectedCode: http.StatusCreated,
			requestBody:  `{"name":"John Doe","lastName":"Johnes","contact":"+5565992526018","address":"Center","brithday":"1996-02-04","cpf":"05000000000","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}`,
			err:          nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.serviceMock = new(service.MockService)
			suite.rabbitMQMock = new(rabbitmq.MockRabbitMQ)

			client := entity.Client{
				Name:     "John Doe",
				LastName: "Johnes",
				Contact:  "+5565992526018",
				Address:  "Center",
				Birthday: "1996-02-04",
				CPF:      "05000000000",
			}

			suite.serviceMock.On("CreateClient", suite.ctx, &client).Return(nil)

			handler := Handler{
				service:     suite.serviceMock,
				environment: &config.Environment{},
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/v1/client", strings.NewReader(cases[key].requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler.CreateClient(suite.ctx)(c)
			suite.Require().NoError(err)
			suite.Require().Equal(cases[key].expectedCode, rec.Code)

			suite.serviceMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *HandlerTestSuite) TestUpdateClient() {
	cases := map[string]struct {
		clientID     string
		requestBody  string
		expectedCode int
		err          error
	}{
		`update client`: {
			clientID:     "647bb6a6-cb08-46f2-a4c7-6d5defa4f573",
			requestBody:  `{"name":"John Doe","lastName":"Johnes","contact":"+5565992526018","address":"Center","brithday":"1996-02-04","cpf":"05000000000","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","deletedAt":null}`,
			expectedCode: http.StatusAccepted,
			err:          nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.serviceMock = new(service.MockService)
			suite.rabbitMQMock = new(rabbitmq.MockRabbitMQ)

			clientID, _ := uuid.Parse(cases[key].clientID)

			client := entity.Client{
				ID:       clientID,
				Name:     "John Doe",
				LastName: "Johnes",
				Contact:  "+5565992526018",
				Address:  "Center",
				Birthday: "1996-02-04",
				CPF:      "05000000000",
			}

			suite.serviceMock.On("UpdateClient", suite.ctx, &client).Return(nil)

			handler := Handler{
				service:     suite.serviceMock,
				environment: &config.Environment{},
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodPut, "/v1/client/"+cases[key].clientID, strings.NewReader(cases[key].requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(cases[key].clientID)

			err := handler.UpdateClient(suite.ctx)(c)
			suite.Require().NoError(err)
			suite.Require().Equal(cases[key].expectedCode, rec.Code)

			suite.serviceMock.AssertExpectations(suite.T())
		})
	}
}

func (suite *HandlerTestSuite) TestDeleteClient() {
	cases := map[string]struct {
		clientID     string
		expectedCode int
		err          error
	}{
		`delete client`: {
			clientID:     "647bb6a6-cb08-46f2-a4c7-6d5defa4f573",
			expectedCode: http.StatusAccepted,
			err:          nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			suite.serviceMock = new(service.MockService)
			suite.rabbitMQMock = new(rabbitmq.MockRabbitMQ)

			clientID, _ := uuid.Parse(cases[key].clientID)

			suite.serviceMock.On("DeleteClient", suite.ctx, clientID).Return(nil)

			handler := Handler{
				service:     suite.serviceMock,
				environment: &config.Environment{},
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, "/v1/client/"+cases[key].clientID, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(cases[key].clientID)

			err := handler.DeleteClient(suite.ctx)(c)
			suite.Require().NoError(err)
			suite.Require().Equal(cases[key].expectedCode, rec.Code)

			suite.serviceMock.AssertExpectations(suite.T())
		})
	}
}
