-- +goose Up
create table note
(
    id         bigserial primary key,
    title      text not null,
    text       text not null,
    author     text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table note;
