package entity

import (
	"time"

	"github.com/google/uuid"
)

type service string

const SERVICE service = "service"

const (
	SERVICE_NAME = "api-test-golang"
	DEV          = "dev"
	PRODUCTION   = "production"
	TEST         = "test"

	CREATE_CLIENT = "create-client"
	UPDATE_CLIENT = "update-client"
)

type Client struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	LastName  string     `json:"lastName" db:"last_name"`
	Contact   string     `json:"contact" db:"contact"`
	Address   string     `json:"address" db:"address"`
	Birthday  string     `json:"brithday" db:"brithday"`
	CPF       string     `json:"cpf" db:"cpf"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

type ResponseWithMeta struct {
	Response
	Pagination *Pagination `json:"metadata"`
}

type Response struct {
	Data interface{} `json:"data"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
	TotalItems int `json:"totalItems"`
}

type QueueRequest struct {
	Param string      `json:"param"`
	Data  interface{} `json:"data"`
}
