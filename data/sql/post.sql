create table post
(
    pid        int auto_increment
        primary key,
    title      varchar(255)                           not null,
    creator    int                                    null,
    content    text                                   null,
    created_at timestamp  default current_timestamp() null,
    last_mod   timestamp                              null,
    hidden     tinyint(1) default 0                   null,
    anonymous  tinyint(1) default 0                   null
);

