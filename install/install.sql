/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.7.20-log : Database - auth
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`auth` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `auth`;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `casbin_rule` */

insert  into `casbin_rule`(`id`,`p_type`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) values (1,'p','日志管理','/logs/log_login:show','*','root','',''),(2,'p','日志管理','/logs/log_login:add','*','root','',''),(3,'p','日志管理','/logs/log_login:edit','*','root','',''),(4,'p','日志管理','/logs/log_login:del','*','root','',''),(5,'p','日志管理','/logs/log_operation:show','*','root','',''),(6,'p','日志管理','/logs/log_operation:add','*','root','',''),(7,'p','日志管理','/logs/log_operation:edit','*','root','',''),(8,'p','日志管理','/logs/log_operation:del','*','root','',''),(9,'p','日志管理','/logs/log_error:show','*','root','',''),(10,'p','日志管理','/logs/log_error:add','*','root','',''),(11,'p','日志管理','/logs/log_error:edit','*','root','',''),(12,'p','日志管理','/logs/log_error:del','*','root','',''),(13,'p','系统设置','/auth-system/menu:show','*','root','',''),(14,'p','系统设置','/auth-system/menu:add','*','root','',''),(15,'p','系统设置','/auth-system/menu:edit','*','root','',''),(16,'p','系统设置','/auth-system/menu:del','*','root','',''),(17,'p','系统设置','/auth-system/domain:show','*','root','',''),(18,'p','系统设置','/auth-system/domain:add','*','root','',''),(19,'p','系统设置','/auth-system/domain:edit','*','root','',''),(20,'p','系统设置','/auth-system/domain:del','*','root','',''),(21,'p','权限管理','/auth-system/menu:show','*','root','',''),(22,'p','权限管理','/auth-system/menu:add','*','root','',''),(23,'p','权限管理','/auth-system/menu:edit','*','root','',''),(24,'p','权限管理','/auth-system/menu:del','*','root','',''),(25,'p','权限管理','/auth-system/domain:show','*','root','',''),(26,'p','权限管理','/auth-system/domain:add','*','root','',''),(27,'p','权限管理','/auth-system/domain:edit','*','root','',''),(28,'p','权限管理','/auth-system/domain:del','*','root','',''),(29,'p','超级管理员','/permission/user:show','*','root','',''),(30,'p','超级管理员','/permission/user:add','*','root','',''),(31,'p','超级管理员','/permission/user:edit','*','root','',''),(32,'p','超级管理员','/permission/user:del','*','root','',''),(33,'p','超级管理员','/permission/dept:show','*','root','',''),(34,'p','超级管理员','/permission/dept:add','*','root','',''),(35,'p','超级管理员','/permission/dept:edit','*','root','',''),(36,'p','超级管理员','/permission/dept:del','*','root','',''),(37,'p','超级管理员','/permission/role:show','*','root','',''),(38,'p','超级管理员','/permission/role:add','*','root','',''),(39,'p','超级管理员','/permission/role:edit','*','root','',''),(40,'p','超级管理员','/permission/role:del','*','root','',''),(41,'p','超级管理员','/auth-system/menu:show','*','root','',''),(42,'p','超级管理员','/auth-system/menu:add','*','root','',''),(43,'p','超级管理员','/auth-system/menu:edit','*','root','',''),(44,'p','超级管理员','/auth-system/menu:del','*','root','',''),(45,'p','超级管理员','/auth-system/domain:show','*','root','',''),(46,'p','超级管理员','/auth-system/domain:add','*','root','',''),(47,'p','超级管理员','/auth-system/domain:edit','*','root','',''),(48,'p','超级管理员','/auth-system/domain:del','*','root','','');

/*Table structure for table `data_perm` */

DROP TABLE IF EXISTS `data_perm`;

CREATE TABLE `data_perm` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '项目域id',
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `perms` varchar(100) NOT NULL DEFAULT '' COMMENT '数据权限标识',
  `perms_rule` varchar(500) NOT NULL DEFAULT '' COMMENT '数据规则',
  `perms_type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '1=分类 2=数据权限',
  `order_num` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '排序',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `remarks` varchar(500) NOT NULL DEFAULT '' COMMENT '说明',
  PRIMARY KEY (`id`),
  KEY `idx_data_perm_domain_id` (`domain_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据权限管理';

/*Data for the table `data_perm` */

/*Table structure for table `department` */

DROP TABLE IF EXISTS `department`;

CREATE TABLE `department` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '部门名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级部门id',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `order_num` int(11) DEFAULT '1' COMMENT '排序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='组织架构表';

/*Data for the table `department` */

insert  into `department`(`id`,`name`,`parent_id`,`create_time`,`update_time`,`order_num`) values (1,'技术部',0,'2018-12-28 00:08:55','2019-04-16 16:38:38',1),(2,'运营部',0,'2018-12-31 06:26:44','2019-03-27 14:58:19',1),(3,'产品部',0,'2019-01-10 09:23:15','2019-03-27 14:58:19',1),(5,'商务部',0,'2019-01-29 19:03:49','2019-03-27 14:58:19',1),(6,'未分配',0,'2019-03-13 17:39:16','2019-03-27 14:58:19',1);

/*Table structure for table `domain` */

DROP TABLE IF EXISTS `domain`;

CREATE TABLE `domain` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `callbackurl` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '回调链接',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '说明',
  `code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标识',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='项目域表';

/*Data for the table `domain` */

insert  into `domain`(`id`,`name`,`callbackurl`,`remark`,`code`,`create_time`,`last_update_time`) values (1,'权限中心','','管理所有后台项目的菜单，权限，鉴权等','root','2018-12-28 16:17:51','2019-03-15 09:51:11'),(2,'测试中心','https://www.baidu.com','此项目用来测试','test','2019-04-16 16:19:38','2019-04-24 17:47:11');

/*Table structure for table `log` */

DROP TABLE IF EXISTS `log`;

CREATE TABLE `log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
  `operation` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户操作',
  `method` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求方法',
  `params` varchar(5000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求参数',
  `time` bigint(20) NOT NULL COMMENT '执行时长(毫秒)',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP地址',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `laste_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='日志表';

/*Data for the table `log` */

/*Table structure for table `menu` */

DROP TABLE IF EXISTS `menu`;

CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级菜单id',
  `domain_id` int(11) NOT NULL DEFAULT '0' COMMENT '项目域id',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `url` tinytext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '路由url',
  `perms` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '授权标识',
  `menu_type` int(11) NOT NULL DEFAULT '0' COMMENT '类型 1=菜单 2=按钮',
  `icon` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
  `order_num` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_menu_domain_id` (`domain_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='菜单表';

/*Data for the table `menu` */

insert  into `menu`(`id`,`parent_id`,`domain_id`,`name`,`url`,`perms`,`menu_type`,`icon`,`order_num`,`create_time`,`last_update_time`) values (1,0,1,'权限管理','/permission','',1,'peoples',1,'2018-12-29 16:15:09','2019-03-20 13:01:46'),(2,1,1,'用户管理','/permission/user','',1,'peoples',1,'2018-12-29 16:15:09','2018-12-29 16:15:09'),(3,2,1,'浏览','','/permission/user:show',2,'',1,'2018-12-29 16:15:09','2018-12-29 16:15:09'),(4,2,1,'添加','','/permission/user:add',2,'',2,'2018-12-29 16:15:09','2018-12-29 16:15:09'),(5,2,1,'修改','','/permission/user:edit',2,'',3,'2018-12-29 16:15:09','2018-12-29 16:15:09'),(6,2,1,'删除','','/permission/user:del',2,'',4,'2018-12-29 16:15:09','2018-12-29 16:15:09'),(7,1,1,'部门管理','/permission/dept','',1,'peoples',2,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(8,7,1,'浏览','','/permission/dept:show',2,'',1,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(9,7,1,'添加','','/permission/dept:add',2,'',2,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(10,7,1,'修改','','/permission/dept:edit',2,'',3,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(11,7,1,'删除','','/permission/dept:del',2,'',4,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(12,1,1,'角色管理','/permission/role','',1,'peoples',3,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(13,12,1,'浏览','','/permission/role:show',2,'',1,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(14,12,1,'添加','','/permission/role:add',2,'',2,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(15,12,1,'修改','','/permission/role:edit',2,'',3,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(16,12,1,'删除','','/permission/role:del',2,'',4,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(17,0,1,'系统设置','/auth-system','',1,'nested',2,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(18,17,1,'菜单管理','/auth-system/menu','',1,'peoples',1,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(19,18,1,'浏览','','/auth-system/menu:show',2,'',1,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(20,18,1,'添加','','/auth-system/menu:add',2,'',2,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(21,18,1,'修改','','/auth-system/menu:edit',2,'',3,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(22,18,1,'删除','','/auth-system/menu:del',2,'',4,'2018-12-29 16:15:10','2018-12-29 16:15:10'),(23,17,1,'项目管理','/auth-system/domain','',1,'peoples',2,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(24,23,1,'浏览','','/auth-system/domain:show',2,'',1,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(25,23,1,'添加','','/auth-system/domain:add',2,'',2,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(26,23,1,'修改','','/auth-system/domain:edit',2,'',3,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(27,23,1,'删除','','/auth-system/domain:del',2,'',4,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(29,28,1,'登录日志','/logs/log_login','',1,'peoples',1,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(30,29,1,'浏览','','/logs/log_login:show',2,'',1,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(31,29,1,'添加','','/logs/log_login:add',2,'',2,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(32,29,1,'修改','','/logs/log_login:edit',2,'',3,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(33,29,1,'删除','','/logs/log_login:del',2,'',4,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(34,28,1,'操作日志','/logs/log_operation','',1,'peoples',2,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(35,34,1,'浏览','','/logs/log_operation:show',2,'',1,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(36,34,1,'添加','','/logs/log_operation:add',2,'',2,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(37,34,1,'修改','','/logs/log_operation:edit',2,'',3,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(38,34,1,'删除','','/logs/log_operation:del',2,'',4,'2018-12-29 16:15:11','2018-12-29 16:15:11'),(39,28,1,'异常日志','/logs/log_error','',1,'peoples',3,'2018-12-29 16:15:12','2018-12-29 16:15:12'),(40,39,1,'浏览','','/logs/log_error:show',2,'',1,'2018-12-29 16:15:12','2018-12-29 16:15:12'),(41,39,1,'添加','','/logs/log_error:add',2,'',2,'2018-12-29 16:15:12','2018-12-29 16:15:12'),(42,39,1,'修改','','/logs/log_error:edit',2,'',3,'2018-12-29 16:15:12','2018-12-29 16:15:12'),(43,39,1,'删除','','/logs/log_error:del',2,'',4,'2018-12-29 16:15:12','2018-12-29 16:15:12'),(66,0,0,'首页','admin/default','',1,'size',1,'2019-01-25 16:34:33','2019-01-25 16:34:33'),(70,69,1,'121','','/permission/role:edit',1,'bug',1,'2019-01-28 10:46:42','2019-01-28 10:46:42');

/*Table structure for table `role` */

DROP TABLE IF EXISTS `role`;

CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `domain_id` int(11) NOT NULL DEFAULT '1' COMMENT '项目域id',
  `role_name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_ids` text COLLATE utf8mb4_unicode_ci,
  `menu_ids_ele` text COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `domain_role` (`domain_id`,`role_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='角色表';

/*Data for the table `role` */

insert  into `role`(`id`,`name`,`domain_id`,`role_name`,`remark`,`menu_ids`,`menu_ids_ele`) values (1,'超级管理员',1,'超级管理员','超级管理员','1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27','1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27'),(2,'系统设置',1,'系统设置','系统设置','17,18,19,20,21,22,23,24,25,26,27','17,18,19,20,21,22,23,24,25,26,27'),(3,'日志管理',1,'日志管理','日志管理','28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43','28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43');

/*Table structure for table `role_data_perm` */

DROP TABLE IF EXISTS `role_data_perm`;

CREATE TABLE `role_data_perm` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `data_perm_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '数据权限id',
  PRIMARY KEY (`id`),
  KEY `idx_rdp_role_id` (`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据权限与角色关系表';

/*Data for the table `role_data_perm` */

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名称',
  `mobile` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `sex` int(11) NOT NULL DEFAULT '0' COMMENT '性别',
  `realname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `department_id` int(11) NOT NULL DEFAULT '0' COMMENT '部门ID',
  `faceicon` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像地址',
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1=正常 2=禁用',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `title` tinytext COLLATE utf8mb4_unicode_ci COMMENT '职位，头衔',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户表';

/*Data for the table `user` */

insert  into `user`(`id`,`username`,`mobile`,`sex`,`realname`,`password`,`salt`,`department_id`,`faceicon`,`email`,`status`,`create_time`,`last_login_time`,`title`) values (1,'wutongci','1862011114',1,'西西','19c13b60ae7ff1a94fc1365e393b3c443b92c10dad438170ddf3ac46fa0e0838d63c1a4e035a69ad240fd8bd6c30329420cf2bccdc5c007ddc7d378b72e1585a','3afff4ba636ce2a1fca2380619b182c16356422faa7157bbe7ee968f40c05bcfe48c61d518ded465c8c7a0adb43c358b3c4a0171d775abad5d88d6b023247dab',1,'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif','lily@hotmail.com',1,'2018-12-22 08:07:59','2018-12-22 08:07:59','developer'),(2,'admin','123123123',1,'admin','7076e523bbc9053eba77d19eb257ee9a4f4dda57061956701034d27cfd5e6ce4c413ad11a73a81e7493f3316ba4854799ff7eac995389029e8d95a7b78093ed1','1d5cb810b443f77c1e1f29317aad408065dc6b708f75f529249240114804835a25885c5be93d8b72e7dc0880a91e52fee160d21a8ffcf9ae55f93a1e2187fa06',1,'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif','111@123.com',1,'2019-02-18 01:30:11','2019-02-18 01:30:11','超管');

/*Table structure for table `user_email` */

DROP TABLE IF EXISTS `user_email`;

CREATE TABLE `user_email` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT '0' COMMENT '用户ID',
  `email` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT 'Email',
  `checked` tinyint(2) DEFAULT '0' COMMENT '是否验证 0 未验证 1 已经验证',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `laste_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户邮件账户表';

/*Data for the table `user_email` */

/*Table structure for table `user_oauth` */

DROP TABLE IF EXISTS `user_oauth`;

CREATE TABLE `user_oauth` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `from` tinyint(3) NOT NULL DEFAULT '1' COMMENT '来源 1.dingtalk 2.wechat 3.facebook',
  `user_id` int(11) DEFAULT NULL COMMENT '用户ID',
  `name` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户姓名',
  `openid` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '第三方openid',
  `unionid` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '第三方unionid',
  `avatar` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '头像',
  `extra` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '扩展信息',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户第三方登陆关联表';

/*Data for the table `user_oauth` */

/*Table structure for table `user_role` */

DROP TABLE IF EXISTS `user_role`;

CREATE TABLE `user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户角色关系表';

/*Data for the table `user_role` */

/*Table structure for table `user_secret` */

DROP TABLE IF EXISTS `user_secret`;

CREATE TABLE `user_secret` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT '0' COMMENT 'userid',
  `account_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '账户名称',
  `secret` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '密钥',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户安全设置-Google 2fa两步验证';

/*Data for the table `user_secret` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
