create table user
(
    uid        int auto_increment
        primary key,
    name       varchar(20)                            not null,
    username   char(10)                               not null,
    email      varchar(255)                           not null,
    password   binary(32)                             not null,
    created_at timestamp  default current_timestamp() not null,
    is_admin   tinyint(1) default 0                   not null,
    banned     tinyint(1) default 0                   not null,
    locked     tinyint(1) default 0                   not null,
    flags      text       default ''                  not null,
    constraint user_username_uindex
        unique (username)
);

