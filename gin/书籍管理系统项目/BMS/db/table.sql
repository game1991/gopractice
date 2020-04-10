---------------------
--这是注释
--BMS
---------------------
create table book (
id bigint(20)  AUTO_INCREMENT PRIMARY KEY,
title varchar(20) NOT NULL ,
price  double(16,2) not null
)engine=InnoDB default charset=utf8mb4;
