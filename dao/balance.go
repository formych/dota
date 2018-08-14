package dao

import (
	"fmt"
	"time"

	"github.com/formych/dota/config"
	"github.com/sirupsen/logrus"
)

// Balance ...
type Balance struct {
	ID        int64     `db:"id"`
	UID       int64     `db:"uid"`
	Balance   int64     `db:"balance"`
	Gbalance  int64     `db:"gbalance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
type balanceDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// BalanceDAO ...
var BalanceDAO = &balanceDAO{
	table:      "balance",
	columns:    "uid, balance, gbalance, created_at, updated_at",
	addColumns: "uid, balance, gbalance, created_at, updated_at",
}

func (bdao *balanceDAO) Add(b *Balance) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($, $2, $3, $4, $5)", bdao.table, bdao.addColumns)
	_, err = config.DB.Exec(exeSQL, b.UID, b.Balance, b.Gbalance, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		logrus.Errorf("dao: insert row failed, sql:[%s], row:[%+v], err:[%s]", exeSQL, *b, err.Error())
		return
	}
	return
}

func (bdao *balanceDAO) Update(uid, balance, gbalance int64) (err error) {
	exeSQL := fmt.Sprintf("UPDATE %s SET balance = %d AND gbalance = %d WHERE uid = %d", bdao.table, balance, gbalance, uid)
	_, err = config.DB.Exec(exeSQL)
	if err != nil {
		logrus.Errorf("dao: insert row failed, sql:[%s], err:[%s]", exeSQL, err.Error())
		return
	}
	return
}

func (bdao *balanceDAO) UpdateBalace(uid, balance int64) (err error) {
	exeSQL := fmt.Sprintf("UPDATE %s SET balance = %d WHERE uid = %d", bdao.table, balance, uid)
	_, err = config.DB.Exec(exeSQL)
	if err != nil {
		logrus.Errorf("dao: insert row failed, sql:[%s], err:[%s]", exeSQL, err.Error())
		return
	}
	return
}

func (bdao *balanceDAO) UpdateGbalance(uid, gbalance int64) (err error) {
	exeSQL := fmt.Sprintf("UPDATE %s SET gbalance = %d WHERE uid = %d", bdao.table, gbalance, uid)
	_, err = config.DB.Exec(exeSQL)
	if err != nil {
		logrus.Errorf("dao: insert row failed, sql:[%s], err:[%s]", exeSQL, err.Error())
		return
	}
	return
}

func (bdao *balanceDAO) Get(uid int64) (res *Balance, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE uid = %d", bdao.table, bdao.columns, uid)
	_, err = config.DB.Exec(exeSQL)
	if err != nil {
		logrus.Errorf("dao: insert row failed, sql:[%s], err:[%s]", exeSQL, err.Error())
		return
	}
	return
}
