-- +goose Up

CREATE TABLE users(
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	updated_at TIMESTAMP NOT NULL DEFAULT now(),
	name varchar(255) NOT NULL
);


-- +goose Down

DROP TABLE users;
