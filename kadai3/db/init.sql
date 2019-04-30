CREATE TABLE users(id serial,name TEXT NOT NULL,email TEXT NOT NULL,created_at TEXT NOT NULL,updated_at TEXT NOT NULL);
ALTER TABLE users ADD PRIMARY KEY(id);