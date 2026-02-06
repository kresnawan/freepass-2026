/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-12.1.2-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: bcc_canteen
-- ------------------------------------------------------
-- Server version	12.1.2-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `account_id` binary(16) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `password` varchar(256) NOT NULL,
  `role` enum('admin','owner','customer') NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  `deactived_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`account_id`) USING BTREE,
  UNIQUE KEY `UNIQUE` (`username`,`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `accounts` VALUES (0x019C262CC97ABF8E7083520CC60EB8BE,'root','root@root.com','','','$argon2id$v=19$m=65536,t=1,p=12$NhqNYBNCS+y8Ucspn0qOfA$fcDNbYg8evXhfQ1HjvdQ4ymW8zudDHjYWKu96sW7Jbo','admin',1,NULL,'2026-02-04 01:03:15','2026-02-04 01:03:15');
INSERT INTO `accounts` VALUES (0x019C32045EC36EF2AF5D836D98E26930,'pakrose','pakrose@gmail.com','Pak','Rose','$argon2id$v=19$m=65536,t=1,p=12$OT6Ns/k68AiymkzVjKYXNg$rsvJH5Xn78DdSosCp/Pzy0n07XNvKOt3gNNj74Dg1i0','owner',1,'2026-02-06 05:16:23','2026-02-06 08:14:33','2026-02-06 12:17:02');
INSERT INTO `accounts` VALUES (0x019C3204DE58C9130417A40EBA6878D9,'burose','burose@gmail.com','Bu','Rose','$argon2id$v=19$m=65536,t=1,p=12$H2XCgcbGVdL/yuXQv58g+A$NYgnL0yNY/fFPeQV8ZjXrJr+Cxi6+hbKjkzfv2o9src','owner',1,NULL,'2026-02-06 08:15:05','2026-02-06 08:15:05');
INSERT INTO `accounts` VALUES (0x019C3205598865E7A561197A8EAEEAFC,'bambang','bambang@gmail.com','Bambang','Pamungkas','$argon2id$v=19$m=65536,t=1,p=12$E22IAkupS0VciljCNVAaDw$hUd0MQnNQ7ljJlKy8GF2ve1BNjeZluRYXjYqzOwtxxE','owner',1,NULL,'2026-02-06 08:15:37','2026-02-06 08:15:37');
INSERT INTO `accounts` VALUES (0x019C3205AED5CEE4B42CDC1DC7B9628B,'tukiran','tukiran@gmail.com','Tukiran','Solikin','$argon2id$v=19$m=65536,t=1,p=12$TFHi4DXdQ3SSIXhXgVpTdw$ViWG/5zXfYBdG7FW2rBI5pjfJVh3Gxf4rirZrnyPH9Y','owner',1,NULL,'2026-02-06 08:15:59','2026-02-06 08:15:59');
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `admin_profile`
--

DROP TABLE IF EXISTS `admin_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin_profile` (
  `account_id` binary(16) NOT NULL,
  PRIMARY KEY (`account_id`) USING BTREE,
  CONSTRAINT `admin_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_profile`
--

LOCK TABLES `admin_profile` WRITE;
/*!40000 ALTER TABLE `admin_profile` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `admin_profile` VALUES (0x019C262CC97ABF8E7083520CC60EB8BE);
/*!40000 ALTER TABLE `admin_profile` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `canteen`
--

DROP TABLE IF EXISTS `canteen`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `canteen` (
  `canteen_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`canteen_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `canteen`
--

LOCK TABLES `canteen` WRITE;
/*!40000 ALTER TABLE `canteen` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `canteen` VALUES (11,'Warung Lalapan Pak Bambang',1);
INSERT INTO `canteen` VALUES (12,'Geprek Kak Rose',1);
INSERT INTO `canteen` VALUES (13,'Kedai Ayam Goreng Tukiran',1);
/*!40000 ALTER TABLE `canteen` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `canteen_ownership`
--

DROP TABLE IF EXISTS `canteen_ownership`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `canteen_ownership` (
  `owner_id` binary(16) NOT NULL,
  `canteen_id` int(11) NOT NULL,
  `owned_at` timestamp NOT NULL DEFAULT current_timestamp(),
  KEY `ownership_owner_id` (`owner_id`),
  KEY `ownership_canteen_id` (`canteen_id`),
  CONSTRAINT `ownership_canteen_id` FOREIGN KEY (`canteen_id`) REFERENCES `canteen` (`canteen_id`),
  CONSTRAINT `ownership_owner_id` FOREIGN KEY (`owner_id`) REFERENCES `owner_profile` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `canteen_ownership`
--

LOCK TABLES `canteen_ownership` WRITE;
/*!40000 ALTER TABLE `canteen_ownership` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `canteen_ownership` VALUES (0x019C3205598865E7A561197A8EAEEAFC,11,'2026-02-06 08:17:48');
INSERT INTO `canteen_ownership` VALUES (0x019C3205AED5CEE4B42CDC1DC7B9628B,13,'2026-02-06 08:18:19');
INSERT INTO `canteen_ownership` VALUES (0x019C3204DE58C9130417A40EBA6878D9,12,'2026-02-06 12:15:19');
INSERT INTO `canteen_ownership` VALUES (0x019C32045EC36EF2AF5D836D98E26930,12,'2026-02-06 12:15:38');
/*!40000 ALTER TABLE `canteen_ownership` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `cart`
--

DROP TABLE IF EXISTS `cart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `cart` (
  `account_id` binary(16) NOT NULL,
  `canteen_id` int(11) NOT NULL,
  `menu_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `price_per_item` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cart`
--

LOCK TABLES `cart` WRITE;
/*!40000 ALTER TABLE `cart` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `cart` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `customer_profile`
--

DROP TABLE IF EXISTS `customer_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer_profile` (
  `account_id` binary(16) NOT NULL,
  `phone_number` varchar(12) NOT NULL DEFAULT '',
  `canteen_points` int(11) NOT NULL DEFAULT 0,
  `instagram` varchar(50) NOT NULL DEFAULT '',
  `bio` varchar(500) NOT NULL DEFAULT '',
  PRIMARY KEY (`account_id`) USING BTREE,
  CONSTRAINT `customer_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer_profile`
--

LOCK TABLES `customer_profile` WRITE;
/*!40000 ALTER TABLE `customer_profile` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `customer_profile` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `feedback`
--

DROP TABLE IF EXISTS `feedback`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `feedback` (
  `feedback_id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` binary(16) NOT NULL,
  `description` varchar(255) NOT NULL,
  `star` enum('1','2','3','4','5') NOT NULL,
  PRIMARY KEY (`feedback_id`),
  UNIQUE KEY `feedback_order_id` (`order_id`) USING BTREE,
  CONSTRAINT `feedback_order_id` FOREIGN KEY (`order_id`) REFERENCES `order` (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feedback`
--

LOCK TABLES `feedback` WRITE;
/*!40000 ALTER TABLE `feedback` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `feedback` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
  `menu_id` int(11) NOT NULL AUTO_INCREMENT,
  `canteen_id` int(11) NOT NULL,
  `menu_name` varchar(50) NOT NULL,
  `price` int(11) NOT NULL,
  `is_removed` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`menu_id`),
  KEY `menu_canteen_id` (`canteen_id`),
  CONSTRAINT `menu_canteen_id` FOREIGN KEY (`canteen_id`) REFERENCES `canteen` (`canteen_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `menu` VALUES (12,12,'Geprek Crispy',14500,0);
INSERT INTO `menu` VALUES (13,12,'Es Teh',3000,0);
INSERT INTO `menu` VALUES (14,11,'Lalapan Bebek',21000,0);
INSERT INTO `menu` VALUES (15,11,'Lalapan Ayam',21000,0);
INSERT INTO `menu` VALUES (16,13,'Nasi Ayam Kalasan',11000,0);
INSERT INTO `menu` VALUES (17,13,'Nasi Ayam Kalasan Jumbo',14000,0);
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `menu_stock`
--

DROP TABLE IF EXISTS `menu_stock`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu_stock` (
  `menu_id` int(11) NOT NULL,
  `stock` int(11) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  KEY `stock_menu_id_fk` (`menu_id`),
  CONSTRAINT `stock_menu_id_fk` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`menu_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu_stock`
--

LOCK TABLES `menu_stock` WRITE;
/*!40000 ALTER TABLE `menu_stock` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `menu_stock` VALUES (12,2,'2026-02-06 08:22:51');
INSERT INTO `menu_stock` VALUES (13,5,'2026-02-06 08:23:05');
INSERT INTO `menu_stock` VALUES (14,7,'2026-02-06 08:27:34');
INSERT INTO `menu_stock` VALUES (15,10,'2026-02-06 08:27:42');
INSERT INTO `menu_stock` VALUES (16,9,'2026-02-06 08:29:41');
INSERT INTO `menu_stock` VALUES (17,3,'2026-02-06 08:29:49');
/*!40000 ALTER TABLE `menu_stock` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `order` (
  `order_id` binary(16) NOT NULL,
  `customer_id` binary(16) NOT NULL,
  `canteen_id` int(11) NOT NULL DEFAULT 0,
  `status` tinyint(1) NOT NULL DEFAULT 6,
  `is_paid` tinyint(1) NOT NULL DEFAULT 0,
  `paid_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`order_id`),
  KEY `order_customer_id_fk` (`customer_id`),
  KEY `order_canteen_id_fk` (`canteen_id`),
  CONSTRAINT `order_canteen_id_fk` FOREIGN KEY (`canteen_id`) REFERENCES `canteen` (`canteen_id`),
  CONSTRAINT `order_customer_id_fk` FOREIGN KEY (`customer_id`) REFERENCES `customer_profile` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order`
--

LOCK TABLES `order` WRITE;
/*!40000 ALTER TABLE `order` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `order` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `order_items`
--

DROP TABLE IF EXISTS `order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_items` (
  `menu_id` int(11) NOT NULL,
  `order_id` binary(16) NOT NULL,
  `quantity` int(11) NOT NULL,
  `price_per_item` int(11) NOT NULL,
  KEY `order_items_menu_id_fk` (`menu_id`),
  KEY `order_items_order_id_fk` (`order_id`),
  CONSTRAINT `order_items_menu_id_fk` FOREIGN KEY (`menu_id`) REFERENCES `menu` (`menu_id`),
  CONSTRAINT `order_items_order_id_fk` FOREIGN KEY (`order_id`) REFERENCES `order` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_items`
--

LOCK TABLES `order_items` WRITE;
/*!40000 ALTER TABLE `order_items` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `order_items` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `owner_profile`
--

DROP TABLE IF EXISTS `owner_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `owner_profile` (
  `account_id` binary(16) NOT NULL,
  `phone_number` varchar(12) NOT NULL DEFAULT '',
  PRIMARY KEY (`account_id`) USING BTREE,
  CONSTRAINT `owner_account_id_fk` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `owner_profile`
--

LOCK TABLES `owner_profile` WRITE;
/*!40000 ALTER TABLE `owner_profile` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `owner_profile` VALUES (0x019C32045EC36EF2AF5D836D98E26930,'');
INSERT INTO `owner_profile` VALUES (0x019C3204DE58C9130417A40EBA6878D9,'');
INSERT INTO `owner_profile` VALUES (0x019C3205598865E7A561197A8EAEEAFC,'');
INSERT INTO `owner_profile` VALUES (0x019C3205AED5CEE4B42CDC1DC7B9628B,'');
/*!40000 ALTER TABLE `owner_profile` ENABLE KEYS */;
UNLOCK TABLES;
commit;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2026-02-06 19:21:40
