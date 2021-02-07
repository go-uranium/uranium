create table comment
(
    cid bpchar(36) not null
        constraint comment_pk
        primary key,
    pid int4 not null
        constraint comment_post_pid_fk
        references post,
    content   text not null,
    markdown  text not null,
    creator int4 not null,
    created_at timestamptz(6) not null,
    last_mod timestamptz(6) not null,
    anonymous bool not null,
    vote_pos _int4(10),
    vote_neg _int4(10) not null
);

