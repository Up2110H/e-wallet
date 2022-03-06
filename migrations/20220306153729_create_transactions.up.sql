CREATE TABLE transactions
(
    id         bigserial      not null primary key,
    wallet_id  bigint         not null
        constraint transactions_wallet_id_foreign
            references wallets
            on update cascade
            on delete cascade,
    amount     numeric(36, 2) not null,
    created_at timestamp(0) default now()
)