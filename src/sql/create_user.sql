use bitcoin;
create table user ( id int primary key not null auto_increment, name varchar(255) not null, email varchar(255) not null unique, password varchar(255) not null, birthdate date not null);