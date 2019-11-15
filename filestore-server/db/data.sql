# 创建数据库fileserver
CREATE DATABASE fileserver DEFAULT CHARACTER SET utf8;

# tbl_file sql
CREATE TABLE `tbl_file` (
`id` INT(11) NOT NULL AUTO_INCREMENT,
`file_sha1` CHAR(40) NOT NULL DEFAULT '',
`file_name` VARCHAR(256) NOT NULL DEFAULT '',
`file_size` BIGINT(20) DEFAULT '0',
`file_addr` VARCHAR(1024) default '',
`create_at` DATETIME DEFAULT NOW(),
`update_at` DATETIME DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP(),
`status` INT(11) NOT NULL DEFAULT '0',
`ext1` INT(11) DEFAULT '0',
`ext2` TEXT,
PRIMARY KEY(`id`),
UNIQUE KEY `idx_file_hash` (`file_sha1`),
KEY `idx_status` (`status`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

# tbl_user sql
 CREATE TABLE `tbl_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT 'encode密码',
  `email` varchar(64) DEFAULT '' COMMENT '邮箱',
  `phone` varchar(128) DEFAULT '' COMMENT 'shouji',
  `email_validated` tinyint(1) DEFAULT '0' COMMENT '邮箱是否验证',
  `phone_validated` tinyint(1) DEFAULT '0' COMMENT '手机是否验证',
  `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `profile` text COMMENT '简历',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等信息)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_phone` (`phone`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

# tbl_user_token sql
CREATE TABLE tbl_user_token (
`id` INT(11) NOT NULL AUTO_INCREMENT,
`user_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户名',
`user_token` CHAR(40) NOT NULL DEFAULT '' COMMENT '用户登录token',
PRIMARY KEY (`id`),
UNIQUE KEY `idx_username` (`user_name`)
)ENGINE = InnoDB DEFAULT CHARSET =utf8;

# tbl_user_file
CREATE TABLE `tbl_user_file` (
`id` INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
`user_name` VARCHAR(64) NOT NULL ,
`file_sha1` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '文件hash',
`file_size` BIGINT(20) DEFAULT '0',
`file_name` VARCHAR(256) NOT NULL DEFAULT  '',
`upload_at` DATETIME DEFAULT CURRENT_TIMESTAMP ,
`last_update` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
`status` INT(11) NOT NULL DEFAULT '0',
KEY `idx_status` (`status`),
KEY `idx_user_id` (`user_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# `last_update` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP 中
# 每次数据库更新, 都会自动更新last_update值