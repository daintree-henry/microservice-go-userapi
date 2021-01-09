CREATE SCHEMA `userdb` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ;

CREATE TABLE `userdb`.`users` (
  `primary_key` BIGINT NOT NULL AUTO_INCREMENT,
  `id` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `phone_number` VARCHAR(45) NOT NULL,
  `status` VARCHAR(45) NULL,
  `created_at` DATETIME NOT NULL,
  `modified_at` DATETIME NOT NULL,
  PRIMARY KEY (`primary_key`),
  UNIQUE INDEX `ID_UNIQUE` (`id` ASC) VISIBLE);

