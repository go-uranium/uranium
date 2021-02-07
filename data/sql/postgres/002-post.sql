create table post
(
    pid int4 default nextval('post_pid_seq'::regclass) not null
        constraint post_pk
        primary key,
    content  text not null,
    markdown text not null
);

