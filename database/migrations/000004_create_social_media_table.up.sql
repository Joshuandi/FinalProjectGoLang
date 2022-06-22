create table social_media(
sm_id SERIAL primary key,
sm_name varchar(50) not null,
sm_url text not null,
user_id int,
constraint user_id foreign key(user_id) references users(u_id) on delete cascade on update cascade
);