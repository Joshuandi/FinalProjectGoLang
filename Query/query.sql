create table users(
u_id SERIAL primary key,
u_username varchar(50) unique not null,
u_email varchar(50) unique not null,
u_pass varchar(70)not null,
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

create table comment(
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
sm_created_at date,
sm_updated_at date,
user_id int,
constraint user_id foreign key(user_id) references users(u_id) on delete cascade on update cascade
);

insert into users (u_username, u_email, u_pass, u_age,u_created_date, u_updated_date)
values ($1, $2, $3, $4, $5, $6);

select*from users;
select*from photo;
select*from comment;
select*from social_media;

drop table social_media ;
drop table comment;
drop table photo;
drop table users;

