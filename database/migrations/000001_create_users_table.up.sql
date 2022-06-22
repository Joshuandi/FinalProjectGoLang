create table users(
u_id SERIAL primary key,
u_username varchar(50) unique not null,
u_email varchar(50) unique not null,
u_pass varchar(70)not null,
u_age int not null,
u_created_date date,
u_updated_date date
);