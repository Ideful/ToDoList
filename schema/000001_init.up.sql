CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE todo_lists (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE users_lists (
    id SERIAL NOT NULL UNIQUE,
    user_id int references users(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);

CREATE TABLE todo_items (
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done boolean not null default false
);

CREATE TABLE lists_items (
    id SERIAL NOT NULL UNIQUE,
    item_id int references todo_items(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);
