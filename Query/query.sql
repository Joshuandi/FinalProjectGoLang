create table users(
u_id SERIAL primary key,
u_username varchar(50) unique not null,
u_email varchar(50) unique not null,
u_pass varchar(50)not null,
u_age int not null,
u_created_date date,
u_updated_date date
);

create table photo(
p_id SERIAL primary key,
p_title text not null,
P_caption text not null,
p_url text not null,
user_id int,
p_created_date date,
p_updated_date date,
constraint user_id foreign key(user_id) references users(u_id) on delete cascade on update cascade
);

create table comment_(
c_id SERIAL primary key,
c_message text not null,
user_id int,
photo_id int,
c_created_date date,
c_updated_date date,
constraint user_id foreign key(user_id) references users(u_id) on delete cascade on update cascade,
constraint photo_id foreign key(photo_id) references photo(p_id) on delete cascade on update cascade
);

create table social_media(
sm_id SERIAL primary key,
sm_name varchar(50) not null,
sm_url text not null,
user_id int,
constraint user_id foreign key(user_id) references users(u_id) on delete cascade on update cascade
);


select*from users;
select*from photo;
select*from comment_;
select*from social_media;

drop table users;
drop table photo;
drop table comment_;
drop table social_media ;