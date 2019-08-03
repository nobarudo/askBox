-- +goose Up
-- +goose StatementBegin
create table question (
    id integer PRIMARY KEY,
    targetUser  varchar(255) NOT NULL,
    question    text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table question;
-- +goose StatementEnd
