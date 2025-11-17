package postgres

import (
	"github.com/go-pg/pg/v10"
)

type LogsRepo struct {
	DB *pg.DB
}

// func (m *LogsRepo) CreateLogs(input model.LogInput) (*model.Log, error) {
// 	Log := &model.Log{

// 	}
// }
