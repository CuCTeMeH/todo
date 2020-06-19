# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.4.11-MariaDB)
# Database: todo
# Generation Time: 2020-06-19 14:41:25 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table lists
# ------------------------------------------------------------

LOCK TABLES `lists` WRITE;
/*!40000 ALTER TABLE `lists` DISABLE KEYS */;

INSERT INTO `lists` (`id`, `uuid`, `user_id`, `name`, `status`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'5008700b-738a-44db-9397-9ca53305777c',1,'test_name','active','2020-06-19 10:17:34','2020-06-19 10:17:34',NULL);

/*!40000 ALTER TABLE `lists` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table tasks
# ------------------------------------------------------------

LOCK TABLES `tasks` WRITE;
/*!40000 ALTER TABLE `tasks` DISABLE KEYS */;

INSERT INTO `tasks` (`id`, `uuid`, `user_id`, `name`, `description`, `status`, `deadline`, `created_at`, `updated_at`, `deleted_at`, `list_id`)
VALUES
	(1,'50fe094d-9f9d-46c6-904d-bda872648540',1,'Test Task Name',NULL,'active','2020-06-19 10:30:29','2020-06-19 10:30:29','2020-06-19 10:30:29',NULL,1),
	(2,'d27bd3d4-5c10-4561-8927-04d1068b3fa8',1,'Test Task Name 2',NULL,'active','2020-06-19 10:30:29','2020-06-19 10:30:29','2020-06-19 10:30:29',NULL,1);

/*!40000 ALTER TABLE `tasks` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `uuid`, `first_name`, `last_name`, `email`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(1,'a8c07e16-1bc7-4554-9131-3e24407e2dbc','First Name','Last Name','email@email.com','2020-06-19 10:31:31','2020-06-19 10:31:31',NULL);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
