/*
 Navicat Premium Data Transfer
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Database       : auth

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : utf-8

 Date: 04/24/2019 16:12:23 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `casbin_rule`
-- ----------------------------
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

-- ----------------------------
--  Table structure for `data_perm`
-- ----------------------------
DROP TABLE IF EXISTS `data_perm`;
CREATE TABLE `data_perm` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `domain_id` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '项目域id',
    `menu_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '菜单ID',
    `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
    `perms` varchar(100) NOT NULL DEFAULT '' COMMENT '数据权限标识',
    `order_num` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '排序',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_data_perm_domain_id` (`domain_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据权限管理';

-- ----------------------------
--  Table structure for `department`
-- ----------------------------
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

-- ----------------------------
--  Table structure for `domain`
-- ----------------------------
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

-- ----------------------------
--  Table structure for `log`
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
     `id` bigint(20) NOT NULL AUTO_INCREMENT,
     `username` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
     `operation` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户操作',
     `method` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求方法',
     `params` varchar(5000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求参数',
     `time` bigint(20) NOT NULL COMMENT '执行时长(毫秒)',
     `ip` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'IP地址',
     `create_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
     `laste_update_time` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='日志表';

-- ----------------------------
--  Table structure for `menu`
-- ----------------------------
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

-- ----------------------------
--  Table structure for `role`
-- ----------------------------
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

-- ----------------------------
--  Table structure for `user`
-- ----------------------------
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
    `last_login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    `title` tinytext COLLATE utf8mb4_unicode_ci COMMENT '职位，头衔',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户表';

-- ----------------------------
--  Table structure for `user_role`
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
   `id` int(11) NOT NULL AUTO_INCREMENT,
   `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
   `role_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
   PRIMARY KEY (`id`),
   KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户角色关系表';

SET FOREIGN_KEY_CHECKS = 1;