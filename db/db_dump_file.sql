-- MySQL dump 10.13  Distrib 8.0.38, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: data_feeds
-- ------------------------------------------------------
-- Server version	9.0.0
use data_feeds;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `agoda_data`
--

DROP TABLE IF EXISTS `agoda_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `agoda_data` (
  `proposal_detail_id` int NOT NULL,
  `proposal_id` int NOT NULL,
  `day_count` int NOT NULL,
  `agoda_hotel_location_id` varchar(50) DEFAULT NULL,
  `agoda_location_name` varchar(255) DEFAULT NULL,
  `start_dt` date DEFAULT NULL,
  `end_dt` date DEFAULT NULL,
  `number_of_rooms` int DEFAULT NULL,
  `number_of_adults` int DEFAULT NULL,
  `number_of_children` int DEFAULT NULL,
  `children_ages_pipe_delimited` varchar(50) DEFAULT NULL,
  `agoda_hotel_name` varchar(255) DEFAULT NULL,
  `agoda_hodel_id` int DEFAULT NULL,
  `agoda_review_score` float DEFAULT NULL,
  `agoda_review_count` int DEFAULT NULL,
  `best_price_per_night` decimal(10,2) DEFAULT NULL,
  `agoda_room_pic_pipe_delimited` text,
  `agoda_hotel_pic_pipe_delimited` text,
  `sleeps_count` int DEFAULT NULL,
  `room_price_per_night` decimal(10,2) DEFAULT NULL,
  `risk_free_booking_binary` tinyint DEFAULT NULL,
  `cancelation_deadline_dt` date DEFAULT NULL,
  `includes_breakfast_bainary` tinyint DEFAULT NULL,
  `includeas_lunch_binary` tinyint DEFAULT NULL,
  `includes_dinner_binary` tinyint DEFAULT NULL,
  `bed_type` varchar(255) DEFAULT NULL,
  `size_sqm` float DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `active_flg` int DEFAULT '1',
  PRIMARY KEY (`proposal_detail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `agoda_data`
--

LOCK TABLES `agoda_data` WRITE;
/*!40000 ALTER TABLE `agoda_data` DISABLE KEYS */;
/*!40000 ALTER TABLE `agoda_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `contact_whatsapp`
--

DROP TABLE IF EXISTS `contact_whatsapp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `contact_whatsapp` (
  `contact_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `surname` varchar(255) DEFAULT NULL,
  `origin_country_id` int DEFAULT NULL,
  `address_1` varchar(255) DEFAULT NULL,
  `address_2` varchar(255) DEFAULT NULL,
  `co_phone_cd` varchar(10) DEFAULT NULL,
  `mobile` varchar(20) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `tag_id` int DEFAULT NULL,
  `agree_contact` int DEFAULT NULL,
  `agree_promo` int DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `active_flg` int DEFAULT '1',
  PRIMARY KEY (`contact_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contact_whatsapp`
--

LOCK TABLES `contact_whatsapp` WRITE;
/*!40000 ALTER TABLE `contact_whatsapp` DISABLE KEYS */;
INSERT INTO `contact_whatsapp` VALUES (1,'Jean','Dupont',1,'123 Rue de Paris','Apt. 4','33','812345678','jean.dupont@example.com',1,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(2,'Marie','Curie',1,'56 Boulevard Saint-Michel','5th Floor','33','812345679','marie.curie@example.com',2,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(3,'Patrick','O\'Connor',2,'15 St. Stephen\'s Green','Suite 3','353','812345680','patrick.oconnor@example.com',3,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(4,'Siobhan','Murphy',2,'47 Merrion Square','Apt. B','353','812345681','siobhan.murphy@example.com',1,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(5,'John','Doe',3,'456 Nyerere Road','Kariakoo','255','812345682','john.doe@example.com',2,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(6,'Asha','Mwinyi',3,'123 Uhuru Street','','','','asha.mwinyi@example.com',2,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(7,'Kamau','Njoroge',4,'','','254','812345684','kamau.njoroge@example.com',2,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(8,'Wanjiku','Mwaniki',4,'234 Kenyatta Avenue','Westlands','254','812345685','wanjiku.mwaniki@example.com',2,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(9,'Pierre','Dupuis',1,'78 Avenue des Champs-?lys?s','7th Floor','33','12345686','pierre.dupuis@example.com',1,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(10,'Aoife','Kelly',2,'23 Grafton Street','Apt. C','353','812345687','aoife.kelly@example.com',4,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(11,'Aoife','Kelly',2,'23 Grafton Street','Apt. C','353','0','aoife.kelly@example.com',3,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(12,'Thomas','Smith',1,'45 Rue de Rivoli','','33','812345688','thomas.smith@example.com',2,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(13,'Fiona','Brown',2,'10 College Green','Suite 5','353','812345689','fiona.brown@example.com',1,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(14,'Grace','Kawawa',3,'25 Bagamoyo Road','Upanga','255','no thanks','grace.kawawa@example.com',1,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(15,'James','Kabiru',4,'12 Kenyatta Avenue','Upper Hill','254','812345691','james.kabiru@example.com',1,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(16,'Emma','Njoroge',4,'99 Moi Avenue','Westlands','254','812345692','emma.njoroge@example.com',1,1,1,'2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1);
/*!40000 ALTER TABLE `contact_whatsapp` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_ref`
--

DROP TABLE IF EXISTS `tag_ref`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag_ref` (
  `tag_id` int NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `active_flg` int DEFAULT '1',
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_ref`
--

LOCK TABLES `tag_ref` WRITE;
/*!40000 ALTER TABLE `tag_ref` DISABLE KEYS */;
INSERT INTO `tag_ref` VALUES (1,'start','2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(2,'middle','2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(3,'comms','2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(4,'end','2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1);
/*!40000 ALTER TABLE `tag_ref` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `second_name` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `active_flg` int DEFAULT '1',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_ref`
--

DROP TABLE IF EXISTS `user_ref`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_ref` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) NOT NULL,
  `second_name` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `active_flg` int DEFAULT '1',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_ref`
--

LOCK TABLES `user_ref` WRITE;
/*!40000 ALTER TABLE `user_ref` DISABLE KEYS */;
INSERT INTO `user_ref` VALUES (1,'Alex','Smith','alex','useralex','2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(2,'Gopal','Patel','gopal','usergopal','2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1),(3,'Engadev','Jones','engadev','userengadev','2024-07-18 15:58:58','2024-07-18 15:58:58',NULL,1);
/*!40000 ALTER TABLE `user_ref` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-07-18 18:01:32
