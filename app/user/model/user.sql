CREATE TABLE `user`
(
    `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
    `nickname`    varchar(255) NOT NULL COMMENT '用户昵称',
    `mobile`      char(11)     NOT NULL COMMENT '手机号',
    `password`    varchar(255) NOT NULL COMMENT '用户密码',
    `gender`      tinyint(3) unsigned NOT NULL DEFAULT '3' COMMENT '性别(1男 2 女 3 保密)',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uni_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户基础信息';