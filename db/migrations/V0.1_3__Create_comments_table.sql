create table comment
(
    id             int         not null primary key auto_increment,
    title          varchar(50) not null,
    body           text        not null,
    published_date datetime    not null,
    user_id        int         not null,
    post_id        int         not null,

    constraint comment_fk_user foreign key (user_id) references `user` (id),
    constraint comment_fk_post foreign key (post_id) references post (id)
);
