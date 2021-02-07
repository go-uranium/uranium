create table "user"
(
    uid int4 default nextval('user_uid_seq'::regclass) not null
        constraint user_pk
        primary key,
    name     varchar(30)  not null,
    username varchar(20)  not null,
    email    varchar(255) not null,
    avatar bpchar(32) not null,
    bio      varchar(511) not null,
    created_at timestamptz(6) not null,
    is_admin bool         not null,
    banned   bool         not null,
    flags _varchar(16)
);

