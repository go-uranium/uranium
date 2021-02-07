create table session
(
    token bpchar(32) not null
        constraint session_pk
        primary key,
    uid int4 not null
        constraint session_user_uid_fk
        references "user",
    ua text        not null,
    ip varchar(39) not null,
    created_at timestamptz(6) not null,
    expire_at timestamptz(6) not null
);

