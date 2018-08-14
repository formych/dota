-- 暂时没有太大业务量, 就不过度设计了
-- (uuid为了以后的分表分库)
-- 当前没有分表分库, 联合查询即可
DROP TABLE IF EXISTS user_auth;
CREATE TABLE IF NOT EXISTS user_auth (
  id serial8, 
  auth_type SMALLINT  NOT NULL,
  identifier VARCHAR(128) NOT NULL,
  certificate VARCHAR(128) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  status SMALLINT NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS uniq_user_auth_ider on user_auth(identifier);
-- CREATE UNIQUE INDEX IF NOT EXISTS uniq_user_auth_uid_type on user_auth(id, auth_type);
CREATE INDEX IF NOT EXISTS idx_user_auth_created on user_auth(created_at);

DROP TABLE IF EXISTS user_balance;
CREATE TABLE IF NOT EXISTS user_balance (
  id serial8 NOT NULL,
  uid BIGINT NOT NULL,
  balance BIGINT NOT NULL DEFAULT 0,
  gbalance BIGINT NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS uniq_user_balance_uid on user_balance(uid);
CREATE INDEX IF NOT EXISTS idx_user_balance_created on user_balance(created_at);


DROP TABLE IF EXISTS user_balance_detail;
CREATE TABLE IF NOT EXISTS user_balance_detail (
  id serial8  NOT NULL ,
  uid BIGINT NOT NULL,
  trade_id VARCHAR(64),
  trade_type SMALLINT NOT NULL,
  pay_type SMALLINT NOT NULL,
  source SMALLINT NOT NULL,
  amount BIGINT NOT NULL,
  balance BIGINT NOT NULL,
  gbalance BIGINT NOT NULL,
  comment VARCHAR(255) NOT NULL DEFAULT '',
  extra_data VARCHAR(255) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS uniq_user_balance_detail_tid on user_balance_detail(trade_id);
CREATE INDEX IF NOT EXISTS idx_user_balance_detail_created on user_balance_detail(created_at);

DROP TABLE IF EXISTS guess_info;
CREATE TABLE IF NOT EXISTS guess_info (
  id serial8  NOT NULL,
  uid BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL,
  settle_type_id BIGINT NOT NULL,
  guess_type_id BIGINT NOT NULL,
  chip_type_id BIGINT NOT NULL,
  fund_pool BIGINT NOT NULL,
  result SMALLINT NOT NULL,
  status SMALLINT NOT NULL,
  info JSON,
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS uniq_guess_info_gid on guess_info(uid);
CREATE INDEX IF NOT EXISTS idx_guess_info_created on guess_info(created_at);

DROP TABLE IF EXISTS guess_record;
CREATE TABLE IF NOT EXISTS guess_record (
  id serial8  NOT NULL,
  uid BIGINT NOT NULL,
  gid BIGINT NOT NULL,
  guess_type_id SMALLINT NOT NULL,
  amount BIGINT NOT NULL,
  result SMALLINT NOT NULL,
  earnings BIGINT NOT NULL,
  odds BIGINT NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_guess_record_created on guess_record(created_at);

DROP TABLE IF EXISTS guess_type;
CREATE TABLE IF NOT EXISTS guess_type (
  id serial8  NOT NULL,
  name VARCHAR(16) NOT NULL,
  "desc" VARCHAR(1024) NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_guess_type_name on guess_type(name);
CREATE INDEX IF NOT EXISTS idx_guess_type_created on guess_type(created_at);

DROP TABLE IF EXISTS settle_type;
CREATE TABLE IF NOT EXISTS settle_type (
  id serial8  NOT NULL,
  name VARCHAR(16) NOT NULL,
  "desc" VARCHAR(1024) NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_settle_type_name on settle_type(name);
CREATE INDEX IF NOT EXISTS idx_settle_type_created on settle_type(created_at);

DROP TABLE IF EXISTS chip_type;
CREATE TABLE IF NOT EXISTS chip_type (
  id serial8  NOT NULL,
  name VARCHAR(16) NOT NULL,
  "desc" VARCHAR(1024) NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_chip_type_name on chip_type(name);
CREATE INDEX IF NOT EXISTS idx_chip_type_created on chip_type(created_at);

DROP TABLE IF EXISTS team;
CREATE TABLE IF NOT EXISTS team (
  id serial8  NOT NULL,
  name VARCHAR(16) NOT NULL,
  info VARCHAR(1024) NOT NULL,
  status SMALLINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS team_name on team(name);
CREATE INDEX IF NOT EXISTS idx_team_created on chip_type(created_at);
