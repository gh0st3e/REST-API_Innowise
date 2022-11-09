CREATE TABLE users(
    id UUID NOT NULL,
    CONSTRAINT id_user PRIMARY KEY(id),
    firstname CHARACTER VARYING(30),
    lastname CHARACTER VARYING(30),
    email CHARACTER VARYING(40) UNIQUE,
    age integer,
    created timestamp
);