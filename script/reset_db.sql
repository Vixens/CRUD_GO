DROP DATABASE IF EXISTS crud_go;
CREATE DATABASE crud_go;

use crud_go;

CREATE TABLE `stores`(
    `StoreID`     int(10) AUTO_INCREMENT PRIMARY KEY,
	`Storename`   varchar(255) DEFAULT NULL,
	`Address`     varchar(255) DEFAULT NULL, 
	`Genre`       varchar(16) DEFAULT NULL,
	`Tel`         varchar(255) DEFAULT NULL,
	`Information` varchar(255) DEFAULT NULL,
	`Tables`      varchar(255) DEFAULT NULL,
	`created_at`  timestamp NULL DEFAULT NULL,
	`updated_at`  timestamp NULL DEFAULT NULL,
	`deleted_at`  timestamp NULL DEFAULT NULL,
	KEY `idx_s_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO stores (Storename,Address,Genre,Tel,Information,Tables) VALUES('hazime','kobe','izakaya','090-5566-5842','tokuninasi',60)