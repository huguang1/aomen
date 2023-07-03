/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Version : 50712
 Source Host           : localhost
 Source Database       : ppgo_api_demo_gin

 Target Server Version : 50712
 File Encoding         : utf-8

 Date: 10/18/2017 22:21:17 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `ppgo_user`
-- ----------------------------
DROP TABLE IF EXISTS `ppgo_user`;
CREATE TABLE `ppgo_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL COMMENT '登录名',
    `password` varchar(64) NOT NULL COMMENT '密码',
    `last_login` datetime(6) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `ppgo_grade`
-- ----------------------------
DROP TABLE IF EXISTS `ppgo_grade`;
CREATE TABLE `ppgo_grade` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `grade` varchar(100) NOT NULL,
    `total_bet` bigint(20) NOT NULL,
    `gold` int(11) NOT NULL,
    `total_gold` int(11) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `ppgo_member`
-- ----------------------------
DROP TABLE IF EXISTS `ppgo_member`;
CREATE TABLE `men_member` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `account` varchar(100) NOT NULL,
    `total_bet` bigint(20) NOT NULL,
    `new_bet` bigint(20) NOT NULL,
    `new_gold` int(11) NOT NULL,
    `total_gold` int(11) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `ppgo_record`
-- ----------------------------
DROP TABLE IF EXISTS `ppgo_record`;
CREATE TABLE `men_record` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `account` varchar(100) NOT NULL,
    `month_amount` bigint(20) NOT NULL,
    `compute` tinyint(1) NOT NULL,
    `date` int(11) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Records of `ppgo_member`
-- ----------------------------
BEGIN;
INSERT INTO `ppgo_member` VALUES ('1', 'haodaquan', '1234'), ('3', 'hell31', 'g2223'), ('4', 'hell31', 'g2223'), ('5', 'hell31', 'g2223'), ('6', 'hell31', 'g2223');
COMMIT;
INSERT INTO `ppgo_user` VALUES ('1', '123', 'python123');
SET FOREIGN_KEY_CHECKS = 1;
