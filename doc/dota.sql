DROP TABLE IF EXISTS `user_auth`;
DROP TABLE IF EXISTS `user_balance`;
DROP TABLE IF EXISTS `user_balance_detail`;
DROP TABLE IF EXISTS `guess_info`;
DROP TABLE IF EXISTS `guess_type`;
DROP TABLE IF EXISTS `guess_record`;
DROP TABLE IF EXISTS `settle_type`;
DROP TABLE IF EXISTS `chip_type`;
// uuid为了以后的分表分库, 暂时没有太大业务量, 就不过度设计了
// 当前没有分表分库, 联合查询即可
CREATE TABLE IF NOT EXISTS `user_auth` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表id, 数据连续性', 
  `uid` BINARY(16) NOT NULL COMMENT '业务用户唯一id, 支持各种uuid',
  `auth_type` TINYINT UNSIGNED NOT NULL COMMENT '1.手机号, 2.邮箱, 3.用户名',
  `identifier` VARCHAR(128) NOT NULL COMMENT '用户标识',
  `certificate` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '用户登录凭证',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', 
  `status` TINYINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ider` (`identifier`),
  UNIQUE KEY `uniq_uid_type` (`uid`, `auth_type`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT="用户授权表";

CREATE TABLE IF NOT EXISTS `user_balance` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表id, 数据连续性',
  `uuid` BINARY(16) NOT NULL COMMENT 'user的唯一标识',
  `balance` BIGINT NOT NULL DEFAULT 0 COMMENT '用户充值余额',
  `gbalance` BIGINT NOT NULL DEFAULT 0 COMMENT '赠送余额, gift balance',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_uid` (`uuid`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户余额表';

CREATE TABLE IF NOT EXISTS `user_balance_detail` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表id, 数据连续性',
  `uuid` BINARY(16) NOT NULL COMMENT 'user唯一标识',
  `trade_id` VARCHAR(50),
  `trade_type` TINYINT NOT NULL COMMENT '交易类型',
  `pay_type` TINYINT NOT NULL COMMENT '支付方式',
  `source` TINYINT NOT NULL COMMENT '交易来源',
  `amount` BIGINT NOT NULL COMMENT '交易金额',
  `balance` BIGINT NOT NULL COMMENT '用户充值余额',
  `gbalance` BIGINT NOT NULL  COMMENT '赠送余额, gift balance',
  `comment` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '交易备注',
  `extra_data` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '额外信息',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_tid` (`trade_id`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户余额详情表';

CREATE TABLE IF NOT EXISTS `guess_info` (
  `id` BIGINT UNSIGNED NOT NULL COMMENT '表id, 数据连续性',
  `guid` BINARY(16) NOT NULL COMMENT '竞猜id',
  `uuid` BINARY(16) NOT NULL COMMENT 'user唯一标识',
  `settle_type_id` BIGINT NOT NULL COMMENT '结算类型',
  `guess_type_id` BIGINT NOT NULL COMMENT '竞猜类型',
  `chip_type_id` BIGINT NOT NULL COMMENT '下注类型',
  `desc` VARCHAR(1024) NOT NULL COMMENT '竞猜描述',
  `fund_pool` BIGINT NOT NULL COMMENT '资金池',
  `status` TINYINT NOT NULL COMMENT '竞猜状态',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` BINARY(16) NOT NULL COMMENT '创建者uuid',
  `updated_by` BINARY(16) NOT NULL COMMENT '更新者uuid',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_guid` (`guid`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户竞猜记录表';

CREATE TABLE IF NOT EXISTS `guess_record` (
  `id` BIGINT UNSIGNED NOT NULL COMMENT '表id, 数据连续性',
  `uuid` BINARY(16) NOT NULL COMMENT 'user唯一标识',
  `guid` BINARY(16) NOT NULL COMMENT 'guess唯一标识',
  `type` TINYINT NOT NULL COMMENT '竞猜类型',
  `amount` BIGINT NOT NULL COMMENT '竞猜金额',
  `result` TINYINT NOT NULL COMMENT '预测结果',
  `earnings` BIGINT NOT NULL COMMENT '竞猜收益',
  `odds` BIGINT NOT NULL COMMENT '赔率',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户参与记录表';

CREATE TABLE IF NOT EXISTS `guess_type` (
  `id` BIGINT UNSIGNED NOT NULL COMMENT '表id, 数据连续性',
  `name` VARCHAR(16) NOT NULL COMMENT '竞猜类型',
  `desc` VARCHAR(1024) NOT NULL COMMENT '竞猜类型描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` BINARY(16) NOT NULL COMMENT '创建者uuid',
  `updated_by` BINARY(16) NOT NULL DEFAULT '更新者uuid',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_type` (`name`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='竞猜类型';

CREATE TABLE IF NOT EXISTS `settle_type` (
  `id` BIGINT UNSIGNED NOT NULL COMMENT '表id, 数据连续性',
  `name` VARCHAR(16) NOT NULL COMMENT '结算类型名字',
  `desc` VARCHAR(1024) NOT NULL COMMENT '结算描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` BINARY(16) NOT NULL COMMENT '创建者uuid',
  `updated_by` BINARY(16) NOT NULL COMMENT '更新者uuid',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_type` (`name`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='结算类型';

CREATE TABLE IF NOT EXISTS `chip_type` (
  `id` BIGINT UNSIGNED NOT NULL COMMENT '表id, 数据连续性',
  `name` VARCHAR(16) NOT NULL COMMENT '筹码类型名字',
  `desc` VARCHAR(1024) NOT NULL COMMENT '竞猜类型描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` BINARY(16) NOT NULL COMMENT '创建者uuid',
  `updated_by` BINARY(16) NOT NULL DEFAULT '更新者uuid',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_type` (`name`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='竞猜类型';

-- CREATE TABLE IF NOT EXISTS ``
INSERT INTO `user_auth` (`uid`, `auth_type`, `identifier`, `certificate`, `created_at`, `update_at`, `status`) VALUES('hhhhhhh', 2, "rrf12", "88888888", "2008-01-01 00:00:00", "2008-01-01 00:00:00", 0);





