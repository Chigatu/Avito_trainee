CREATE TABLE banner 
(
	id serial not null unique,
	name varchar(255) not null,
	username varchar(255) not null unique,
	password_hash varchar(255) not null
);

CREATE TABLE tag
(
	id serial not null unique,
	title varchar(255) not null,
	description varchar(255) 
);

CREATE TABLE feature
(
	id serial not null unique,
	user_id int references users (id) on delete cascade  not null, 
	list_id int references todo_lists (id) on delete cascade  not null
);