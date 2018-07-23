package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/formych/dota/config"
	"github.com/sirupsen/logrus"
)

// UserAuth 用户认证表
type UserAuth struct {
	ID          int64     `db:"id"`
	AuthType    int8      `db:"auth_type"`
	Identifier  string    `db:"identifier"`
	Certificate string    `db:"certificate"`
	CreatedAt   time.Time `db:"created_at"`
	UpdateAt    time.Time `db:"updated_at"`
	Status      int8      `db:"status"`
}

// UserAuthDAO ...
type userAuthDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// UserAuthDAO ...
var UserAuthDAO = userAuthDAO{
	table:      "user_auth",
	columns:    "id, auth_type, identifier, certificate, created_at, updated_at, status",
	addColumns: "auth_type, identifier, certificate, created_at, updated_at, status",
}

// Add 增加一条记录
func (udao *userAuthDAO) Add(u *UserAuth) (id int64, err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4, $5, $6)", udao.table, udao.addColumns)
	result, err := config.DB.Exec(exeSQL, u.AuthType, u.Identifier, u.Certificate, u.CreatedAt, u.UpdateAt, u.Status)
	if err != nil {
		logrus.Errorf("[dao]: insert record failed, sql:[%s], user_auth:[%+v], err:[%s]", exeSQL, *u, err.Error())
		return
	}
	if id, err = result.RowsAffected(); err != nil {
		logrus.Errorf("[dao]: get affected rows failed, sql:[%s], user_auth:[%+v], err:[%s]", exeSQL, *u, err.Error())
	}
	return
}

// Get 通过唯一标识获取
func (udao *userAuthDAO) Get(identifier string) (userAuth *UserAuth, err error) {
	userAuth = &UserAuth{}
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE identifier = $1", udao.columns, udao.table)
	err = config.DB.Get(userAuth, exeSQL, identifier)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logrus.Errorf("[dao]: get record by identifier failed, sql:[%s], identifier:[%s], err:[%s]", exeSQL, identifier, err.Error())
	}
	return
}
