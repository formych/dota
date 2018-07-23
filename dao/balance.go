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
	addColumns string
}

// BalanceDAO ...
var BalanceDAO = &balanceDAO{
	table:      "balance",
	addColumns: "uid, balance, gbalance, created_at, updated_at",
}

func (bdao *balanceDAO) Add(b *Balance) (id int64, err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($, $2, $3, $4, $5)", bdao.table, bdao.addColumns)
	res, err := config.DB.Exec(exeSQL, b.UID, b.Balance, b.Gbalance, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		logrus.Errorf("dao: insert row failed, sql:[%s], row:[%+v], err:[%s]", exeSQL, *b, err.Error())
		return
	}
	id, err = res.RowsAffected()
	if err != nil {
		logrus.Errorf("dao: get rows affected failed, sql:[%s], row:[%+v], err:[%s]", exeSQL, *b, err.Error())
	}
	return
}
