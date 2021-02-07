create table post_info
(
    pid int4 not null
        constraint post_info_pk
        primary key
        constraint post_info_post_pid_fk
        references post,
    title     varchar(255) not null,
    creator int4 not null,
    created_at timestamptz(6) not null,
    last_mod timestamptz(6) not null,
    replies int4 not null,
    views int4 not null,
    activity timestamptz(6) not null,
    hidden    bool         not null,
    anonymous bool         not null,
    vote_pos _int4(10),
    vote_neg _int4(10)
);

