-- +goose Up
-- +goose StatementBegin
create table user (
    id integer PRIMARY KEY,
    name varchar(255) NOT NULL,
    nickName varchar(255) NOT NULL,
    pass varchar(255) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
