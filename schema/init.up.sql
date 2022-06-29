CREATE TABLE links
(
    id serial PRIMARY KEY,
    short_link varchar(10) not null unique,
    full_link varchar(255) not null unique
);

CREATE INDEX ON links(short_link);
CREATE INDEX ON links(full_link);