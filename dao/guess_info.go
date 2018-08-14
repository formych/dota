package dao

import (
	"fmt"
	"time"

	"github.com/formych/dota/config"
	"github.com/sirupsen/logrus"
)

// GuessInfo 竞猜信息表
type GuessInfo struct {
	ID           int64     `json:"id,omitempty" form:"id" db:"id"`
	UID          int64     `json:"uid" form:"uid" db:"uid"`
	Name         string    `json:"name" form:"name" db:"name" `
	SettleTypeID int64     `json:"settle_type_id" form:"settle_type_id" db:"settle_type_id" `
	GuessTypeID  int64     `json:"guess_type_id" form:"guess_type_id" db:"chip_type_id" `
	ChipTypeID   int64     `json:"chip_type_id" form:"chip_type_id" db:"chip_type_id" `
	Result       int64     `json:"result" form:"result" db:"result" `
	FundPool     int64     `json:"fund_pool" form:"fund_pool" db:"fund_pool"`
	Info         string    `json:"Info" form:"Info" db:"Info" `
	StartTime    time.Time `json:"start_time" form:"start_time" db:"start_time" `
	EndTime      time.Time `json:"end_time" form:"end_time" db:"end_time" `
	CreatedAt    time.Time `json:"created_at" form:"created_at" db:"created_at"`
	UpdateAt     time.Time `json:"updated_at" form:"updated_at" db:"updated_at"`
	Status       int8      `json:"status" form:"updated_at" db:"status"`
}

// Option ...
type Option struct {
	Name  string `json:"name" form:"name" db:"name"`
	Value int    `json:"value" form:"value" db:"name"`
}

// Info ...
type Info struct {
	Teams   []int64
	Options []*Option
}

// guessInfoDAO ...
type guessInfoDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// GuessInfoDAO ...
var GuessInfoDAO = &guessInfoDAO{
	table:      "guess_info",
	columns:    "id, uid, settle_type_id, guess_type_id, chip_type_id, result, fund_pool, \"desc\", created_at, updated_at, status",
	addColumns: "uid, name, settle_type_id, guess_type_id, chip_type_id, result, fund_pool, info, start_time, end_time, created_at, updated_at, status",
}

// Add 增加一条记录
func (gdao *guessInfoDAO) Add(g *GuessInfo) (id int64, err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id", gdao.table, gdao.addColumns)
	err = config.DB.QueryRow(exeSQL, g.UID, g.Name, g.SettleTypeID, g.GuessTypeID, g.ChipTypeID, g.Result, g.FundPool, g.Info, g.StartTime, g.EndTime, g.CreatedAt, g.UpdateAt, g.Status).Scan(&id)
	if err != nil {
		logrus.Errorf("[dao](: insert record failed, sql:[%s], GuessInfo:[%+v], err:[%s]", exeSQL, *g, err.Error())
		return
	}
	return
}

// Delete 删除一条记录
func (gdao *guessInfoDAO) Delete(gid, uid int64) (err error) {
	exeSQL := fmt.Sprintf("Update %s SET status = %d WHERE id = %d and uid = %d", gdao.table, -1, gid, uid)
	_, err = config.DB.Exec(exeSQL)
	if err != nil {
		logrus.Errorf("[dao](: delete record failed, sql:[%s], err:[%s]", exeSQL, err.Error())
	}
	return
}

// Update 更新一条记录
// func (gdao *guessInfoDAO) Update(g *GuessInfo) (err error) {
// 	exeSQL := fmt.Sprintf("UPDATE %s SET uid = $1, settle_type_id = $2, guess_type_id = $3, chip_type_id = $4, result = $5, fund_pool = $6, \"desc\" = $7, updated_at = $8, statu = $9 WHERE id = %d)", gdao.table, g.ID)
// 	_, err = config.DB.Exec(exeSQL, g.SettleTypeID, g.GuessTypeID, g.ChipTypeID, g.Result, g.FundPool, g.Result, g.UpdateAt, g.Status)
// 	if err != nil {
// 		logrus.Errorf("[dao](: insert record failed, sql:[%s], GuessInfo:[%+v], err:[%s]", exeSQL, *g, err.Error())
// 		return
// 	}
// 	return
// }

// // Get 根据id获取
// func (gdao *guessInfoDAO) Get(id int64) (res []*GuessInfo, err error) {
// 	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE id = %d", gdao.columns, gdao.table, id)
// 	err = config.DB.Get(&res, exeSQL)
// 	if err == sql.ErrNoRows {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
// 	}
// 	return
// }

// // Get 根据uid获取
// func (gdao *guessInfoDAO) GetByUID(uid int64) (res []*GuessInfo, err error) {
// 	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE uid = %d", gdao.columns, gdao.table, uid)
// 	err = config.DB.Select(&res, exeSQL)
// 	if err == sql.ErrNoRows {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
// 	}
// 	return
// }

// // Get 获取部分记录
// func (gdao *guessInfoDAO) GetSub(start, limit int64) (res []*GuessInfo, err error) {
// 	exeSQL := fmt.Sprintf("SELECT %s FROM %s limit %d, %d", gdao.columns, gdao.table, (start-1)*limit, limit)
// 	err = config.DB.Select(&res, exeSQL)
// 	if err == sql.ErrNoRows {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
// 	}
// 	return
// }

// // Get 获取全部
// func (gdao *guessInfoDAO) GetAll() (guessRecords []*GuessInfo, err error) {
// 	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE 'status' = 1", gdao.columns, gdao.table)
// 	err = config.DB.Select(guessRecords, exeSQL)
// 	if err == sql.ErrNoRows {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
// 	}
// 	return
// }

// // Count 统计总数量
// func (gdao *guessInfoDAO) Count() (total int64, err error) {
// 	exeSQL := fmt.Sprintf("SELECT count(1) FROM %s", gdao.table)
// 	err = config.DB.Get(&total, exeSQL)
// 	if err == sql.ErrNoRows {
// 		return 0, nil
// 	}
// 	if err != nil {
// 		logrus.Errorf("[dao]: get all records by failed, sql:[%s], err:[%s]", exeSQL, err.Error())
// 	}
// 	return
// }
