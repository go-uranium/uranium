create table sign_up
(
    token bpchar(36) not null
        constraint sign_up_pk
        primary key,
    email varchar(255) not null,
    created_at timestamptz(6) not null,
    expire_at timestamptz(6) not null
);

