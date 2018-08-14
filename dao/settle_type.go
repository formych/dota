package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/formych/dota/config"
	"github.com/sirupsen/logrus"
)

// SettleType 筹码类型表
type SettleType struct {
	ID        int64     `json:"id" form:"id" db:"id"`
	Name      string    `josn:"name" form:"name" db:"desc"`
	Desc      string    `json:"desc" form:"desc" db:"desc"`
	CreatedAt time.Time `json:"created_at" form:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"updated_at" form:"updated_at" db:"updated_at"`
	Status    int8      `json:"status" form:"status" db:"status"`
}

// settleTypeDAO ...
type settleTypeDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// SettleTypeDAO ...
var SettleTypeDAO = &settleTypeDAO{
	table:      "chip_type",
	columns:    "id, name, \"desc\", created_at, updated_at, status",
	addColumns: "name, \"desc\", created_at, updated_at, status",
}

// Add 增加一条记录
func (sdao *settleTypeDAO) Add(s *SettleType) (err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4, $5)", sdao.table, sdao.addColumns)
	_, err = config.DB.Exec(exeSQL, s.Name, s.Desc, s.CreatedAt, s.UpdateAt, s.Status)
	if err != nil {
		logrus.Errorf("[dao]: insert record failed, sql:[%s], SettleType:[%+v], err:[%s]", exeSQL, *s, err.Error())
		return
	}
	return
}

// Get 获取
func (sdao *settleTypeDAO) GetAll() (guessTypes []*SettleType, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE 'status' = 1", sdao.columns, sdao.table)
	err = config.DB.Select(&guessTypes, exeSQL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
	}
	return
}
