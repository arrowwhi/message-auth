-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id bigserial,
    name varchar(50),
    email varchar(50) unique,
    password varchar(100),
    created_at timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
