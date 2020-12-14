create table sign_up
(
    email      varchar(255)                                             not null,
    token      char(36)  default uuid()                                 not null,
    created_at timestamp default current_timestamp()                    null,
    expire_at  timestamp default (current_timestamp() + interval 1 day) null,
    constraint sign_up_email_uindex
        unique (email)
);

