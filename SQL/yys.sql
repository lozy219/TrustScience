/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : yys

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 06/09/2018 17:23:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for base_server
-- ----------------------------
DROP TABLE IF EXISTS `base_server`;
CREATE TABLE `base_server`  (
  `server_code` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '服务器区服code',
  `server_name` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '服务器区服名字',
  PRIMARY KEY (`server_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for base_shikigami
-- ----------------------------
DROP TABLE IF EXISTS `base_shikigami`;
CREATE TABLE `base_shikigami`  (
  `shikigami_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '式神code',
  `shikigami_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '式神名字',
  `shikigami_win` int(10) NULL DEFAULT NULL COMMENT '式神胜场',
  `shikigami_lose` int(10) NULL DEFAULT NULL COMMENT '式神败场',
  `weight_coefficient` double(20, 0) NULL DEFAULT NULL COMMENT '式神权重系数',
  PRIMARY KEY (`shikigami_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of base_shikigami
-- ----------------------------
INSERT INTO `base_shikigami` VALUES ('1', '茨木童子', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('10', '阎魔', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('100', '百目鬼', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('101', '日和坊', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('102', '惠比寿', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('103', '烟烟罗', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('104', '犬夜叉', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('105', '杀生丸', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('106', '鬼切', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('11', '一目连', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('12', '御行达摩', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('13', '荒', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('14', '辉夜姬', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('15', '彼岸花', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('16', '雪童子', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('17', '奴良陆生', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('18', '山风', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('19', '玉藻前', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('2', '大天狗', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('20', '御馔津', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('21', '卖药郎', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('22', '鬼灯', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('23', '追月神', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('24', '桃花妖', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('25', '九命猫', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('26', '鲤鱼精', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('27', '三尾狐', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('28', '雪女', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('29', '座敷童子', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('3', '荒川之主', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('30', '饿鬼', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('31', '鸦天狗', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('32', '巫蛊师', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('33', '鬼使白', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('34', '鬼使黑', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('35', '河童', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('36', '狸猫', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('37', '孟婆', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('38', '犬神', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('39', '童男', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('4', '花鸟卷', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('40', '童女', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('41', '跳跳弟弟', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('42', '跳跳妹妹', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('43', '雨女', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('44', '食发鬼', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('45', '武士之灵', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('46', '骨女', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('47', '兵俑', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('48', '丑时之女', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('49', '独眼小僧', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('5', '酒吞童子', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('50', '管狐', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('51', '鬼女红叶', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('52', '椒图', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('53', '山兔', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('54', '跳跳哥哥', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('55', '铁鼠', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('56', '萤草', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('57', '蝴蝶精', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('58', '傀儡师', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('59', '凤凰火', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('6', '两面佛', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('60', '海坊主', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('61', '觉', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('62', '判官', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('63', '青蛙瓷器', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('64', '山童', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('65', '首无', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('66', '吸血姬', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('67', '妖狐', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('68', '食梦貘', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('69', '妖琴师', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('7', '青行灯', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('70', '清姬', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('71', '镰鼬', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('72', '姑获鸟', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('73', '二口女', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('74', '白狼', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('75', '青坊主', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('76', '古笼火', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('77', '万年竹', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('78', '夜叉', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('79', '樱花妖', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('8', '小鹿男', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('80', '般若', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('81', '络新妇', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('82', '黑童子', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('83', '白童子', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('84', '金鱼姬', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('85', '鸩', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('86', '以津真天', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('87', '熏', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('88', '数珠', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('89', '小袖之手', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('9', '妖刀姬', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('90', '弈', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('91', '虫师', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('92', '蜜桃', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('93', '猫掌柜', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('94', '阿香', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('95', '面灵气', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('96', '匣中少女', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('97', '小松丸', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('98', '兔丸', NULL, NULL, NULL);
INSERT INTO `base_shikigami` VALUES ('99', '书翁', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for field_record
-- ----------------------------
DROP TABLE IF EXISTS `field_record`;
CREATE TABLE `field_record`  (
  `field_code` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '记录场次code',
  `server_code` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '服务器区服code',
  `server_name` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '服务器区服名字',
  `shikigami_code_1` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_1` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_2` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_2` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_3` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_3` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_4` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_4` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_5` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_5` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_6` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_6` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_7` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_7` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_8` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_8` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_9` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_9` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_code_10` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `shikigami_name_10` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`field_code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;