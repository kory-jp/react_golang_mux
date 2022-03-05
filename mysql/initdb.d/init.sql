create table if not exists users (
	id integer primary key auto_increment,
	name varchar(50) NOT NULL,
	email varchar(50) NOT NULL UNIQUE,
	password varchar(60) NOT NULL,
	created_at datetime default current_timestamp
);

create table if not exists todos (
	id integer primary key auto_increment,
	user_id integer NOT NULL,
	title varchar(50) NOT NULL,
	content text NOT NULL,
	image_path varchar(100),
	isFinished boolean NOT NULL,
	created_at datetime default current_timestamp
);