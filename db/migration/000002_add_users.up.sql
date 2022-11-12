create table users (
                       username varchar primary key ,
                       password varchar not null ,
                       full_name varchar not null ,
                       email varchar unique  not null ,
                       password_changed_at timestamptz not null default '0001-01-01 00:00:00Z',
                       created_at timestamptz not null default now()
);

alter table accounts
    add constraint accounts__owner_fk
        foreign key (owner) references users (username);

alter table accounts
    add constraint accounts__unique_owner_currency
    unique (owner, currency);
