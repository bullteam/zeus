/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.7.14 : Database - casbin
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`casbin` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `casbin`;

/*Table structure for table `casbin_rule` */

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `p_type` varchar(255) NOT NULL DEFAULT '',
  `v0` varchar(255) NOT NULL DEFAULT '',
  `v1` varchar(255) NOT NULL DEFAULT '',
  `v2` varchar(255) NOT NULL DEFAULT '',
  `v3` varchar(255) NOT NULL DEFAULT '',
  `v4` varchar(255) NOT NULL DEFAULT '',
  `v5` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='casbin权限表';

/*Data for the table `casbin_rule` */

insert  into `casbin_rule`(`id`,`p_type`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) values (1,'p','日志管理','/logs/log_login:show','*','root','',''),(2,'p','日志管理','/logs/log_login:add','*','root','',''),(3,'p','日志管理','/logs/log_login:edit','*','root','',''),(4,'p','日志管理','/logs/log_login:del','*','root','',''),(5,'p','日志管理','/logs/log_operation:show','*','root','',''),(6,'p','日志管理','/logs/log_operation:add','*','root','',''),(7,'p','日志管理','/logs/log_operation:edit','*','root','',''),(8,'p','日志管理','/logs/log_operation:del','*','root','',''),(9,'p','日志管理','/logs/log_error:show','*','root','',''),(10,'p','日志管理','/logs/log_error:add','*','root','',''),(11,'p','日志管理','/logs/log_error:edit','*','root','',''),(12,'p','日志管理','/logs/log_error:del','*','root','',''),(13,'p','系统设置','/auth-system/menu:show','*','root','',''),(14,'p','系统设置','/auth-system/menu:add','*','root','',''),(15,'p','系统设置','/auth-system/menu:edit','*','root','',''),(16,'p','系统设置','/auth-system/menu:del','*','root','',''),(17,'p','系统设置','/auth-system/domain:show','*','root','',''),(18,'p','系统设置','/auth-system/domain:add','*','root','',''),(19,'p','系统设置','/auth-system/domain:edit','*','root','',''),(20,'p','系统设置','/auth-system/domain:del','*','root','',''),(21,'p','权限管理','/auth-system/menu:show','*','root','',''),(22,'p','权限管理','/auth-system/menu:add','*','root','',''),(23,'p','权限管理','/auth-system/menu:edit','*','root','',''),(24,'p','权限管理','/auth-system/menu:del','*','root','',''),(25,'p','权限管理','/auth-system/domain:show','*','root','',''),(26,'p','权限管理','/auth-system/domain:add','*','root','',''),(27,'p','权限管理','/auth-system/domain:edit','*','root','',''),(28,'p','权限管理','/auth-system/domain:del','*','root','',''),(29,'p','超级管理员','/permission/user:show','*','root','',''),(30,'p','超级管理员','/permission/user:add','*','root','',''),(31,'p','超级管理员','/permission/user:edit','*','root','',''),(32,'p','超级管理员','/permission/user:del','*','root','',''),(33,'p','超级管理员','/permission/dept:show','*','root','',''),(34,'p','超级管理员','/permission/dept:add','*','root','',''),(35,'p','超级管理员','/permission/dept:edit','*','root','',''),(36,'p','超级管理员','/permission/dept:del','*','root','',''),(37,'p','超级管理员','/permission/role:show','*','root','',''),(38,'p','超级管理员','/permission/role:add','*','root','',''),(39,'p','超级管理员','/permission/role:edit','*','root','',''),(40,'p','超级管理员','/permission/role:del','*','root','',''),(41,'p','超级管理员','/auth-system/menu:show','*','root','',''),(42,'p','超级管理员','/auth-system/menu:add','*','root','',''),(43,'p','超级管理员','/auth-system/menu:edit','*','root','',''),(44,'p','超级管理员','/auth-system/menu:del','*','root','',''),(45,'p','超级管理员','/auth-system/domain:show','*','root','',''),(46,'p','超级管理员','/auth-system/domain:add','*','root','',''),(47,'p','超级管理员','/auth-system/domain:edit','*','root','',''),(48,'p','超级管理员','/auth-system/domain:del','*','root','','');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
