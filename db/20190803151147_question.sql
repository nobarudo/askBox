-- +goose Up
-- +goose StatementBegin
create table question (
    id integer PRIMARY KEY,
    targetUser  varchar(255) NOT NULL,
    text    text NOT NULL,
    reply   text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table question;
-- +goose StatementEnd
