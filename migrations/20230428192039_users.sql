-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
      id            serial              constraint users_pk PRIMARY KEY,
      email         varchar(100)        unique,
      username      varchar(50),
      password      varchar(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
