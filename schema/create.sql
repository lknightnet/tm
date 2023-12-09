create table projects(
    iduser varchar not null ,
    id varchar not null primary key,
    name varchar not null,
    description varchar
);

create table notes(
    id varchar not null primary key,
    idproject varchar not null,
    description varchar,
    completeness boolean not null ,
    FOREIGN KEY(idproject) REFERENCES projects(id)
);

create table projectusers(
    iduser varchar not null,
    idproject varchar not null
);