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