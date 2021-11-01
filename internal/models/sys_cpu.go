package models

import (
	"fmt"
	"go-sys/internal/pkg/database"
	"gopkg.in/guregu/null.v4"
)

type SysCpuBlock struct {
	DateTime null.Time   `db:"date_time"`
	HostName null.String `db:"host_name"`
	Sys      null.Int    `db:"sys"`
	Idle     null.Int    `db:"idle"`
	Usr      null.Int    `db:"usr"`
	Wait     null.Int    `db:"wait"`
	Total    null.Int    `db:"total"`
}

func (s SysCpuBlock) Insert(tx *database.SQLTX, item *SysCpuBlock) (err error) {
	query := "INSERT INTO sys_cpu VALUES "
	result, err := tx.Exec(query)
	if err != nil {
		panic(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}

	return
}
