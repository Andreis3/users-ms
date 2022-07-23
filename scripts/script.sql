create table users
(
    id          text not null
        constraint table_name_pk
            primary key
        constraint table_name_id_uindex
            unique,
    user_name   text,
    password    text,
    first_name  text,
    last_name   text,
    email       text,
    cpf         text,
    created_at  date,
    modified_at date
);

alter table users
    owner to root;


