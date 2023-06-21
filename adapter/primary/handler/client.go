package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

var ErrInvalidParameters = "invalid parameters"

// @Summary List clients
// @Description Retrieves a list of clients
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} entity.ResponseWithMeta
// @Router /v1/client [get]
func (h Handler) ListClient(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			page = 1
		}

		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			limit = 10
		}

		clients, pagination, err := h.service.ListClient(ctx, page, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Bad Request")
		}

		response := entity.ResponseWithMeta{
			Response: entity.Response{
				Data: clients,
			},
			Pagination: pagination,
		}

		return c.JSON(http.StatusOK, response)
	}
}

// @Summary Get a client
// @Description Retrieves a client by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "Client ID"
// @Success 200 {object} entity.Client
// @Router /v1/client/{id} [get]
func (h Handler) GetClient(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrInvalidParameters)
		}

		client, err := h.service.GetClient(ctx, clientID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, entity.Response{
			Data: client,
		})
	}
}

// CreateClient godoc
// @Summary Create a client
// @Description Creates a new client
// @Tags Users
// @Accept json
// @Produce json
// @Success 201
// @Router /v1/client [post]
func (h Handler) CreateClient(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		client := &entity.Client{}

		err := c.Bind(client)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrInvalidParameters)
		}

		err = h.service.CreateClient(ctx, client)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, nil)
	}
}

// @Summary Update a client
// @Description Updates an existing client
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "Client ID"
// @Success 202
// @Router /v1/client/{id} [put]
func (h Handler) UpdateClient(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrInvalidParameters)
		}

		client := &entity.Client{}

		err = c.Bind(client)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrInvalidParameters)
		}

		client.ID = clientID

		err = h.service.UpdateClient(ctx, client)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusAccepted, nil)
	}
}

// @Summary Delete a client
// @Description Deletes a client by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "Client ID"
// @Success 202
// @Router /v1/client/{id} [delete]
func (h Handler) DeleteClient(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "invalid parameters")
		}

		err = h.service.DeleteClient(ctx, clientID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error deleting client")
		}

		return c.JSON(http.StatusAccepted, nil)
	}
}
