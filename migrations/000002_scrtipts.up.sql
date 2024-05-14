CREATE TABLE scripts (
                           id serial not null PRIMARY KEY,
                           name varchar not null unique,
                           result varchar
);