use bitcoin;
create table bitcoin(id int primary key not null auto_increment, quantity int not null, total mysql> create table bitcoin(id int primary key not null aut int not null, total double not null, date datetime default now(), type varchar(255), person_id int);

ALTER TABLE `bitcoin`.`bitcoin` 
ADD INDEX `fk_bitcoin_1_idx` (`person_id` ASC);
ALTER TABLE `bitcoin`.`bitcoin` 
ADD CONSTRAINT `fk_bitcoin_1`
  FOREIGN KEY (`person_id`)
  REFERENCES `bitcoin`.`user` (`id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION;
