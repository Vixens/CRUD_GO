DROP DATABASE IF EXISTS crud_go;
CREATE DATABASE crud_go;

use crud_go;

CREATE TABLE `stores`(
    `id`     int(10) AUTO_INCREMENT PRIMARY KEY,
	`created_at`  timestamp NULL DEFAULT NULL,
	`updated_at`  timestamp NULL DEFAULT NULL,
	`deleted_at`  timestamp NULL DEFAULT NULL,
	`storename`   varchar(255) DEFAULT NULL,
	`loc` 	      varchar(255) DEFAULT NULL, 
	`genre`       varchar(255) DEFAULT NULL,
	`tel`         varchar(255) DEFAULT NULL,
	`information` varchar(255) DEFAULT NULL,
	`tables`      varchar(255) DEFAULT NULL,

	KEY `idx_s_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO stores (storename,loc,genre,tel,information,tables) VALUES('hazime','kobe','izakaya','090-5566-5842','tokuninasi',60)