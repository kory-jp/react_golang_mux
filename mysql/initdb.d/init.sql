USE react_golang_mux;

CREATE TABLE IF NOT EXISTS users (
	id integer PRIMARY KEY AUTO_INCREMENT,
	name varchar(19) NOT NULL,
	email varchar(29) NOT NULL UNIQUE,
	password varchar(60) NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS todos (
	id integer PRIMARY KEY AUTO_INCREMENT,
	user_id integer NOT NULL,
	title varchar(49) NOT NULL,
	content text NOT NULL,
	image_path varchar(300),
	isFinished boolean NOT NULL,
	importance integer NOT NULL,
	urgency integer NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
	INDEX usr_ind  (user_id),
	FOREIGN KEY (user_id)
		REFERENCES users(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tags (
	id integer PRIMARY KEY AUTO_INCREMENT,
	value varchar(50) NOT NULL,
	label varchar(50) NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_tag_relations (
	id integer PRIMARY KEY AUTO_INCREMENT,
	todo_id integer NOT NULL,
	tag_id integer NOT NULL,
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
	INDEX tod_ind (todo_id),
	FOREIGN KEY (todo_id)
		REFERENCES todos(id)
		ON DELETE CASCADE,
	INDEX tag_ind (tag_id),
	FOREIGN KEY (tag_id)
		REFERENCES tags(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS task_cards (
	id integer PRIMARY KEY AUTO_INCREMENT,
	user_id integer NOT NULL,
  todo_id integer NOT NULL,
  title varchar(49) NOT NULL,
  purpose text,
  content text,
  memo text,
  isFinished boolean NOT NULL,
  created_at datetime DEFAULT CURRENT_TIMESTAMP,
  INDEX td_ind (todo_id),
  FOREIGN KEY (todo_id)
		REFERENCES todos(id)
    ON DELETE CASCADE,
	INDEX usr_ind (user_id),
	FOREIGN KEY (user_id)
		REFERENCES users(id)
		ON DELETE CASCADE
);