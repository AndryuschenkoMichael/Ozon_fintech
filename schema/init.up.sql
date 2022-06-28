CREATE TABLE links
(
    id serial PRIMARY KEY,
    short_link varchar(10) not null unique,
    full_link varchar(255) not null unique
);