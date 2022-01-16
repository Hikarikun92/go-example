create table `user`
(
    id       int          not null primary key auto_increment,
    username varchar(255) not null
);

create table user_credentials
(
    user_id  int          not null primary key,
    password varchar(255) not null,

    constraint user_credentials_fk_user foreign key (user_id) references `user` (id)
);

create table user_roles
(
    user_id int          not null,
    roles   varchar(255) not null,

    constraint user_roles_fk_user foreign key (user_id) references `user` (id)
);