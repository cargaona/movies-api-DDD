create database movies; 
create table movies(id varchar primary key, name varchar not null, director_name varchar not null, director_last_name varchar not null, year int not null, created timestamp not null);
"""
CREATE USER movies WITH SUPERUSER PASSWORD 'example';
"""
