create table accounts (
                          id bigserial primary key ,
                          owner varchar not null ,
                          balance bigint not null ,
                          currency varchar not null default 'AZN',
                          created_at timestamptz not null default now()
);

create table entries (
                         id bigserial primary key ,
                         account_id bigint references accounts(id) not null ,
                         amount bigint not null,
                         created_at timestamptz not null default now()
);

create table transfers (
                           id bigserial primary key ,
                           from_account_id bigint references accounts(id) not null ,
                           to_account_id bigint references accounts(id) not null ,
                           amount bigint not null ,
                           created_at timestamptz not null default now()
);