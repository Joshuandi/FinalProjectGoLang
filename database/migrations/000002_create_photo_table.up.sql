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
