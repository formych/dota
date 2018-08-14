package dao

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/formych/dota/config"
	"github.com/sirupsen/logrus"
)

// ChipType 筹码类型表
type ChipType struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Desc      string    `db:"desc"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
	Status    int8      `db:"status"`
}

// chipTypeDAO ...
type chipTypeDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// ChipTypeDAO ...
var ChipTypeDAO = &chipTypeDAO{
	table:      "chip_type",
	columns:    "id, name, \"desc\", created_at, updated_at, status",
	addColumns: "name, \"desc\", created_at, updated_at, status",
}

// Add 增加一条记录
func (cdao *chipTypeDAO) Add(c *ChipType) (id int64, err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4, $5)", cdao.table, cdao.addColumns)
	result, err := config.DB.Exec(exeSQL, c.Name, c.Desc, c.CreatedAt, c.UpdateAt, c.Status)
	if err != nil {
		logrus.Errorf("[dao]: insert record failed, sql:[%s], ChipType:[%+v], err:[%s]", exeSQL, *c, err.Error())
		return
	}
	if id, err = result.RowsAffected(); err != nil {
		logrus.Errorf("[dao]: get affected rows failed, sql:[%s], ChipType:[%+v], err:[%s]", exeSQL, *c, err.Error())
	}
	return
}

// Get 通过唯一标识获取
func (cdao *chipTypeDAO) Get(name string) (chipType *ChipType, err error) {
	chipType = new(ChipType)
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE status = 1", cdao.columns, cdao.table)
	err = config.DB.Get(chipType, exeSQL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
	}
	return
}

// GetAll 获取所有
func (cdao *chipTypeDAO) GetAll() (chipTypes []*ChipType, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE status = 1", cdao.columns, cdao.table)
	err = config.DB.Select(&chipTypes, exeSQL)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
	}
	return
}
