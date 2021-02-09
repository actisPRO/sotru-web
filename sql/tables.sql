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

create table web_ips
(
    user_id varchar(255) not null,
    ip varchar(255) not null,
    last_used datetime not null,
    constraint web_ips_web_users_id_fk
        foreign key (user_id) references web_users (id)
            on update cascade on delete cascade
);

create table web_xboxes
(
    user_id varchar(255) not null,
    xbox text not null,
    last_used datetime not null,
    constraint web_xboxes_web_users_id_fk
        foreign key (user_id) references web_users (id)
            on update cascade on delete cascade
);
