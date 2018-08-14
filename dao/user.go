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
	ID          int64     `form:"id" json:"id,omitempty" db:"id"`
	AuthType    int8      `form:"auth_type" json:"auth_type" db:"auth_type"`
	Identifier  string    `form:"identifier" json:"identifier" db:"identifier" binding:"required"`
	Certificate string    `form:"certificate" json:"certificate" db:"certificate" binding:"required"`
	CreatedAt   time.Time `form:"created_at" json:"created_at" db:"created_at"`
	UpdateAt    time.Time `form:"updated_at" json:"updated_at" db:"updated_at"`
	Status      int8      `form:"status" json:"status" db:"status"`
}

// UserAuthDAO ...
type userAuthDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// UserAuthDAO ...
var UserAuthDAO = &userAuthDAO{
	table:      "user_auth",
	columns:    "id, auth_type, identifier, certificate, created_at, updated_at, status",
	addColumns: "auth_type, identifier, certificate, created_at, updated_at, status",
}

// Add 增加一条记录
func (udao *userAuthDAO) Add(u *UserAuth) (id int64, err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", udao.table, udao.addColumns)
	err = config.DB.QueryRowx(exeSQL, u.AuthType, u.Identifier, u.Certificate, u.CreatedAt, u.UpdateAt, u.Status).Scan(&id)
	if err != nil {
		logrus.Errorf("[dao]: insert record failed, sql:[%s], user_auth:[%+v], err:[%s]", exeSQL, *u, err.Error())
		return
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

func (udao *userAuthDAO) UpdatePassword(uid int64, passHash string) (err error) {
	exeSQL := fmt.Sprintf("UPDATE %s SET certificate = %s WHERE uid = %d", udao.table, passHash, uid)
	_, err = config.DB.Exec(exeSQL)
	if err != nil {
		logrus.Errorf("[dao]: update user password failed, sql:[%s], identifier:[%s], err:[%s]", exeSQL, passHash, err.Error())
	}
	return
}
