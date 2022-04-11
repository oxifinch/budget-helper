DROP TABLE IF EXISTS user;

-- Defining tables
CREATE TABLE user (
    user_id integer not null unique primary key autoincrement,
    username text not null unique,
    password text not null
);

-- Example data: users
INSERT INTO user (username, password) VALUES (
    "joseph",
    "beans"
);

INSERT INTO user (username, password) VALUES (
    "jean-paul",
    "paris"
);
