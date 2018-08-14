package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/formych/dota/config"
	"github.com/sirupsen/logrus"
)

// GuessRecord 筹码类型表
type GuessRecord struct {
	ID        int64     `db:"id"`
	UID       int64     `db:"uid"`
	GID       int64     `db:"gid"`
	Amount    int64     `db:"amount"`
	Result    int64     `db:"result"`
	Earnings  int64     `db:"earnings"`
	Odds      int64     `db:"odds"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"created_at"`
	Status    int8      `db:"status"`
}

// guessRecordDAO ...
type guessRecordDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// GuessRecordDAO ...
var GuessRecordDAO = &guessRecordDAO{
	table:      "guess_record",
	columns:    "id, uid, gid, amount, result, earnings, odds, created_at, updated_at, 'status'",
	addColumns: "uid, gid, amount, result, earnings, odds, created_at, updated_at, 'status'",
}

// Add 增加一条记录
func (gdao *guessRecordDAO) Add(g *GuessRecord) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)", gdao.table, gdao.addColumns)
	_, err = config.DB.Exec(exeSQL, g.UID, g.GID, g.Amount, g.Result, g.Earnings, g.Odds, g.CreatedAt, g.UpdateAt, g.Status)
	if err != nil {
		logrus.Errorf("[dao](: insert record failed, sql:[%s], GuessRecord:[%+v], err:[%s]", exeSQL, *g, err.Error())
		return
	}
	return
}

// Get 通过唯一标识获取
func (gdao *guessRecordDAO) GetAll() (guessRecords []*GuessRecord, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE 'status' = 1", gdao.columns, gdao.table)
	err = config.DB.Get(guessRecords, exeSQL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
	}
	return
}
