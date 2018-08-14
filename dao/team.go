package dao

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/formych/dota/config"
)

// Team _
type Team struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Info      string    `json:"info" db:"info"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Status    int8      `json:"status" db:"status"`
}

type teamDAO struct {
	db         string
	table      string
	columns    string
	addColumns string
}

// TeamDAO _
var TeamDAO = &teamDAO{
	table:      "team",
	columns:    "id, name, info, created_at, updated_at, status",
	addColumns: "name, info, created_at, updated_at, status",
}

func (d *teamDAO) Add(t *Team) (id int64, err error) {
	exeSQL := fmt.Sprintf("INSERT INTO %s (%s) VALUES($1, $2, $3, $4, $5) RETURN", d.table, d.addColumns)

	err = config.DB.QueryRow(exeSQL, t.Name, t.Info, t.CreatedAt, t.UpdatedAt, t.Status).Scan(id)
	if err != nil {
		logrus.Errorf("Insert into %s failed, team:[%+v]err:[%s]", t, err.Error())
	}
	return

}

func (d *teamDAO) GetAll() (teams []*Team, err error) {
	exeSQL := fmt.Sprintf("SELECT %s FROM %s WHERE status = 1", d.columns, d.table)
	err = config.DB.Select(&teams, exeSQL)
	if err != nil {
		logrus.Errorf("Get all records failed, sql:[%s], err:[%s]", exeSQL, err.Error())
	}
	return
}
