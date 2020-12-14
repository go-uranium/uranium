create table comment
(
    cid        int auto_increment
        primary key,
    pid        int                                    not null,
    creator    int                                    not null,
    content    text                                   not null,
    created_at timestamp  default current_timestamp() not null,
    anonymous  tinyint(1) default 0                   null
);

