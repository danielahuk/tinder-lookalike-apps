# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.4.27-MariaDB)
# Database: tinder_db
# Generation Time: 2024-05-18 17:45:31 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table features
# ------------------------------------------------------------

DROP TABLE IF EXISTS `features`;

CREATE TABLE `features` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `value` varchar(255) NOT NULL,
  `status` int(1) DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `features` WRITE;
/*!40000 ALTER TABLE `features` DISABLE KEYS */;

INSERT INTO `features` (`id`, `name`, `price`, `value`, `status`, `created_at`)
VALUES
	(1,'Extend Quota',10000.00,'10',1,'2024-05-17 15:08:44'),
	(2,'Premium Member',15000.00,'1',1,'2024-05-17 15:08:44');

/*!40000 ALTER TABLE `features` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table members
# ------------------------------------------------------------

DROP TABLE IF EXISTS `members`;

CREATE TABLE `members` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `gender` varchar(50) NOT NULL,
  `label` varchar(50) NOT NULL,
  `quota` int(11) DEFAULT 10,
  `status` int(1) DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;

INSERT INTO `members` (`id`, `name`, `email`, `password`, `gender`, `label`, `quota`, `status`, `created_at`)
VALUES
	(1,'Member M 1','memberm1@gmail.com','$2a$10$9p7izkdaorse4bYysUVwGuU.tsDSMC6RU.9Nf8fGUdELlHMGa1gyy','Male','Premium',99,1,'2024-05-17 20:29:10'),
	(2,'Member F 1','memberf1@gmail.com','$2a$10$i6FZHR7I/k1JQ9NNiV4gPeO/k3Ga7.9IYDK8mNEz8I6frWQy6uViW','Female','Regular',10,1,'2024-05-17 20:56:38'),
	(3,'Member M 2','memberm2@gmail.com','$2a$10$XUa.onUSOQqfRpWUEYDXmOzogOYM945ghxKgUK.GjTSt4Z/cASnsu','Male','Regular',10,1,'2024-05-17 23:07:29'),
	(4,'Member F 2','memberf2@gmail.com','$2a$10$hDl.L1q0reZmOpJFtyF9TuY6CUjWk151i7TqxIWimvLUMswIZn6Pu','Female','Regular',10,1,'2024-05-18 12:44:24');

/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table partnership
# ------------------------------------------------------------

DROP TABLE IF EXISTS `partnership`;

CREATE TABLE `partnership` (
  `member_id1` int(11) DEFAULT NULL,
  `member_id2` int(11) DEFAULT NULL,
  `status` int(1) DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  KEY `member_id1` (`member_id1`),
  KEY `member_id2` (`member_id2`),
  CONSTRAINT `partnership_ibfk_1` FOREIGN KEY (`member_id1`) REFERENCES `members` (`id`),
  CONSTRAINT `partnership_ibfk_2` FOREIGN KEY (`member_id2`) REFERENCES `members` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `partnership` WRITE;
/*!40000 ALTER TABLE `partnership` DISABLE KEYS */;

INSERT INTO `partnership` (`member_id1`, `member_id2`, `status`, `created_at`)
VALUES
	(1,2,2,'2024-05-18 21:35:03'),
	(1,4,1,'2024-05-18 22:00:01');

/*!40000 ALTER TABLE `partnership` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table transaction_details
# ------------------------------------------------------------

DROP TABLE IF EXISTS `transaction_details`;

CREATE TABLE `transaction_details` (
  `transaction_id` int(11) DEFAULT NULL,
  `feature_id` int(11) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `status` int(1) DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  KEY `transaction_id` (`transaction_id`),
  KEY `feature_id` (`feature_id`),
  CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`),
  CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`feature_id`) REFERENCES `features` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `transaction_details` WRITE;
/*!40000 ALTER TABLE `transaction_details` DISABLE KEYS */;

INSERT INTO `transaction_details` (`transaction_id`, `feature_id`, `qty`, `price`, `status`, `created_at`)
VALUES
	(1,1,1,10000.00,1,'2024-05-18 20:42:01'),
	(2,2,1,15000.00,1,'2024-05-18 20:42:06');

/*!40000 ALTER TABLE `transaction_details` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table transactions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `status` int(1) DEFAULT 1,
  `member_id` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `member_id` (`member_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`member_id`) REFERENCES `members` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;

INSERT INTO `transactions` (`id`, `date`, `status`, `member_id`, `created_at`)
VALUES
	(1,'2024-05-18',1,1,'2024-05-18 20:42:01'),
	(2,'2024-05-18',1,1,'2024-05-18 20:42:06');

/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
