-- sqlite ver 3

----------
-- USER --
----------
-- insert once
CREATE TABLE "t_user" (
	"id"	        INTEGER NOT NULL,
	"uname"	        TEXT NOT NULL UNIQUE,
	"created_at"	TIMESTAMP NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);

-- insert once, can be updated
CREATE TABLE "t_user_detail" (
	"id"	        INTEGER NOT NULL,
	"id_user"	INTEGER NOT NULL,
	"first_name"	TEXT NOT NULL,
	"last_name"	TEXT NOT NULL DEFAULT '',
	"email"	        TEXT NOT NULL DEFAULT '',
	"password"	TEXT NOT NULL,
	"flags"	        INTEGER NOT NULL DEFAULT 0,
	"created_at"	TIMESTAMP NOT NULL,
	"updated_at"	TIMESTAMP,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("id_user") REFERENCES "t_user"("id")
);

-- always insert
CREATE TABLE "t_user_detail_history" (
	"id"	                INTEGER NOT NULL,
	"id_user_detail"	INTEGER NOT NULL,
	"first_name"	        TEXT NOT NULL,
	"last_name"	        TEXT NOT NULL DEFAULT '',
	"email"	                TEXT NOT NULL DEFAULT '',
	"password"	        TEXT NOT NULL,
	"flags"	                INTEGER NOT NULL DEFAULT 0,
	"created_at"	        TIMESTAMP NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("id_user_detail") REFERENCES "t_user_detail"("id")
);


----------
-- AUTH --
----------
/*
 * flags: 0 -> logged in
 *        1 -> logged out
 */
-- always insert, can be updated
CREATE TABLE "t_auth" (
	"id"	        INTEGER NOT NULL,
	"token"	        TEXT UNIQUE NOT NULL,
	"flags"         INTEGER NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT)
);


----------
-- TODO --
----------
-- insert once, can be updated
CREATE TABLE "t_todo" (
	"id"	                INTEGER NOT NULL,
	"id_user"	        INTEGER NOT NULL,
	"label"	                TEXT NOT NULL,
	"is_active"		INTEGER NOT NULL,
	"created_at"	        TIMESTAMP NOT NULL,
	"updated_at"	        TIMESTAMP,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("id_user") REFERENCES "t_user"("id")
);

-- always insert
CREATE TABLE "t_todo_history" (
	"id"	                INTEGER NOT NULL,
	"id_todo"	        INTEGER NOT NULL,
	"label"	                TEXT NOT NULL,
	"is_active"		INTEGER NOT NULL,
	"created_at"	        TIMESTAMP NOT NULL,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("id_todo") REFERENCES "t_todo"("id")
);

-- insert once, can be updated
CREATE TABLE "t_todo_item" (
	"id"	                INTEGER NOT NULL,
	"id_todo"	        INTEGER NOT NULL,
	"title"	                TEXT NOT NULL,
	"description"	        TEXT NOT NULL,
	"deadline"		TIMESTAMP,
        "status"                INTEGER NOT NULL DEFAULT 0,
	"created_at"	        TIMESTAMP NOT NULL,
	"updated_at"	        TIMESTAMP,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("id_todo") REFERENCES "t_todo"("id")
);

-- always insert
CREATE TABLE "t_todo_item_history" (
	"id"	                INTEGER NOT NULL,
	"id_todo_item"	        INTEGER NOT NULL,
	"title"	                TEXT NOT NULL,
	"description"	        TEXT NOT NULL,
	"deadline"		TIMESTAMP,
        "status"                INTEGER NOT NULL DEFAULT 0,
	"created_at"	        TIMESTAMP,
	PRIMARY KEY("id" AUTOINCREMENT),
	FOREIGN KEY("id_todo_item") REFERENCES "t_todo_item"("id")
);


----------
-- BLOG --
----------

