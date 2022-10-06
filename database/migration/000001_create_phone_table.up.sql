DROP TABLE IF EXISTS `phone`;

CREATE TABLE `phone`
(
    `id`    bigint unsigned NOT NULL AUTO_INCREMENT,
    `name`  varchar(255) NOT NULL,
    `brand` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;