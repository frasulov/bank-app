alter table if exists accounts
drop constraint if exists  accounts__unique_owner_currency;

alter table if exists accounts
drop constraint if exists accounts__owner_fk ;

drop table if exists users;