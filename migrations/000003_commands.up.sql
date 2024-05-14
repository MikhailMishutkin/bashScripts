CREATE TABLE commands (
                            id serial not null PRIMARY KEY,
                            command varchar not null ,
                            scr_id integer,
                            foreign key (scr_id) REFERENCES scripts (id)
);