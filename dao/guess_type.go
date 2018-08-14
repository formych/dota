package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/formych/dota/config"
	"github.com/sirupsen/logrus"
)

// GuessType 筹码类型表
type GuessType struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Desc      string    `db:"desc"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"created_at"`
	Status    int8      `db:"status"`
}

// guessTypeDAO ...
type guessTypeDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// GuessTypeDAO ...
var GuessTypeDAO = &guessTypeDAO{
	table:      "chip_type",
	columns:    "id, name, 'desc', created_at, updated_at, 'status'",
	addColumns: "name, 'desc', created_at, updated_at, 'status'",
}

// Add 增加一条记录
func (gdao *guessTypeDAO) Add(g *GuessType) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4, $5)", gdao.table, gdao.addColumns)
	_, err = config.DB.Exec(exeSQL, g.Name, g.Desc, g.CreatedAt, g.UpdateAt, g.Status)
	if err != nil {
		logrus.Errorf("[dao]: insert record failed, sql:[%s], GuessType:[%+v], err:[%s]", exeSQL, *g, err.Error())
		return
	}
	return
}

// Get 通过唯一标识获取
func (gdao *guessTypeDAO) GetAll() (guessTypes []*GuessType, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE 'status' = 1", gdao.columns, gdao.table)
	err = config.DB.Select(&guessTypes, exeSQL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
	}
	return
}
