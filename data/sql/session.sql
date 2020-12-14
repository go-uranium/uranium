create table session
(
    token     char(36)  default uuid()                                  not null
        primary key,
    uid       int                                                       not null,
    UA        varchar(255)                                              null,
    IP        varchar(39)                                               null,
    time      timestamp default current_timestamp()                     null,
    expire_at timestamp default (current_timestamp() + interval 30 day) null
);

