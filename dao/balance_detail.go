package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/formych/dota/config"
	"github.com/sirupsen/logrus"
)

// BalanceDetail ...
type BalanceDetail struct {
	ID        int64     `db:"id"`
	UID       int64     `db:"uid"`
	TradeID   string    `db:"trade_id"`
	TradeType int8      `db:"trade_type"`
	PayType   int8      `db:"pay_type"`
	Source    int8      `db:"source"`
	Amount    int64     `db:"amount"`
	Balance   int64     `db:"balance"`
	Gbalance  int64     `db:"gbalance"`
	Comment   string    `db:"comment"`
	ExtraData string    `db:"extra_data"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// BalanceDetailDAO ...
type balanceDetailDAO struct {
	db         string
	table      string
	addColumns string
}

// BalanceDetailDAO ...
var BalanceDetailDAO = &balanceDetailDAO{db: "", table: "user_balance",
	addColumns: "uuid, trade_id, trade_type, pay_type, source, amount, balance, gbalance, comment, ext_data, created_at, updated_at"}

// Add 增加一条余额变动记录
func (bdao *balanceDetailDAO) Add(b *BalanceDetail) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)", bdao.table, bdao.addColumns)
	_, err = config.DB.Exec(exeSQL, b.UID, b.TradeID, b.TradeType, b.PayType, b.Source, b.Amount, b.Balance, b.Gbalance, b.Comment, b.ExtraData, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		logrus.Errorf("dao: insert row failed, sql:[%s], row:[%+v], err:[%s]", exeSQL, *b, err.Error())
		return
	}
	return
}

func (bdao *balanceDetailDAO) GetAll(uid int64) (res []*BalanceDetail, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM WHERE uid = %d", bdao.table, uid)
	err = config.DB.Get(&res, exeSQL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logrus.Errorf("dao: get rows failed, sql:[%s], err:[%s]", exeSQL, err.Error())
	}
	return
}
