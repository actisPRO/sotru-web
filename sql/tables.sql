create table web_users
(
    id varchar(255) not null,
    username text null,
    registered_on datetime null,
    last_login datetime null,
    avatar_url text null,
    access_token text null,
    refresh_token text null,
    access_expiration datetime null,
    constraint web_users_pk
        primary key (id)
);