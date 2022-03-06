CREATE TABLE users
(
    id            bigserial not null primary key,
    email         varchar   not null unique,
    key           varchar   not null,
    identified_at timestamp(0)
)