DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS budget;

---- Defining tables ----
CREATE TABLE user (
    user_id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE budget ( 
    budget_id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
    start_date TEXT NOT NULL,
    end_date TEXT NOT NULL,
    owner_id INTEGER NOT NULL,
    FOREIGN KEY(owner_id) REFERENCES user(user_id)
    
);
-------------------------

-- Example data: users
INSERT INTO user (username, password) VALUES (
    "joseph",
    "beans"
);

INSERT INTO user (username, password) VALUES (
    "jean-paul",
    "paris"
);

-- Example data: budgets
INSERT INTO budget(start_date, end_date, owner_id) VALUES (
    "2022-04-25",
    "2022-05-25",
    1
);
