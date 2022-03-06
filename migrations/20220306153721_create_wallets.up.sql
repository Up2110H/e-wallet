CREATE TABLE wallets
(
    id      bigserial      not null primary key,
    user_id bigint         not null
        constraint wallets_user_id_foreign
            references users
            on update cascade
            on delete cascade,
    amount  numeric(36, 2) not null
)