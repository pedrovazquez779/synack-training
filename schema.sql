drop table if exists tags_videos;
drop table if exists tags;
drop table if exists videos;
drop table if exists users;

create table users
(
    id   integer      not null primary key auto_increment,
    name varchar(100) not null
);

create table videos
(
    id      integer      not null primary key auto_increment,
    user_id integer      not null,
    name    varchar(100) not null,
    deleted bit          not null default false,

    foreign key (user_id) references users (id)
);

create table tags
(
    id      integer      not null primary key auto_increment,
    name    varchar(100) not null,
    deleted bit          not null default false
);

create table tags_videos
(
    video_id int not null,
    tag_id   int not null,

    primary key (video_id, tag_id),
    foreign key (video_id) references videos (id),
    foreign key (tag_id) references tags (id)
);

insert into users (name)
values ('john');
insert into users (name)
values ('wayne');

insert into videos (user_id, name) values (1, 'Batman Begins');
insert into videos (user_id, name) values (1, 'Dark Knight');
insert into videos (user_id, name) values (1, 'Dark Knight Rises');
insert into videos (user_id, name) values (1, 'X-Men');
insert into videos (user_id, name) values (1, 'Matrix');
insert into videos (user_id, name) values (2, 'Star Wars');
insert into videos (user_id, name) values (2, 'Kung Fu Panda');

insert into tags (name, deleted) values ('Action', false);
insert into tags (name, deleted) values ('Drama', false);

insert into tags_videos (video_id, tag_id) values (1, 1);
insert into tags_videos (video_id, tag_id) values (1, 2);
insert into tags_videos (video_id, tag_id) values (2, 2);