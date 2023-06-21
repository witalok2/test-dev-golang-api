-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "client"
(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" character varying NOT NULL,
  last_name character varying DEFAULT NULL,
  contact character varying(20) DEFAULT NULL,
  "address" character varying NOT NULL,
  birthday character varying(10) NOT NULL,
  cpf   character varying(12),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "client";
-- +goose StatementEnd
