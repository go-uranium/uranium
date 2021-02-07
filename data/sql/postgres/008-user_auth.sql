create table user_auth
(
    uid int4 not null
        constraint user_auth_pk
        primary key
        constraint user_auth_user_uid_fk
        references "user",
    password bpchar(32) not null,
    locked         bool         not null,
    security_email varchar(255) not null
);

