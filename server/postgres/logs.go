package postgres

import (
	"time"

	"github.com/SBPH-Matthew/testosterone-tracker/dbmodels"
	"github.com/SBPH-Matthew/testosterone-tracker/graph/model"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type LogsRepo struct {
	DB *pg.DB
}

// func (m *LogsRepo) CreateLogs(input model.LogInput) (*model.Log, error) {
// 	Log := &model.Log{

// 	}
// }

func (m *LogsRepo) CreateLogs(input model.LogInput, userID string) (*model.Log, error) {
	notes := ""
	if input.Notes != nil {
		notes = *input.Notes
	}

	logs := &dbmodels.Log{
		ID:              uuid.New().String(),
		UserID:          userID,
		LogDate:         time.Now().Format(time.RFC3339),
		EnergyLevel:     int(input.EnergyLevel),
		Mood:            int(input.Mood),
		Libido:          int(input.Libido),
		SleepHours:      input.SleepHours,
		ExerciseMinutes: input.ExerciseMinutes,
		StressLevel:     int(input.StressLevel),
		Notes:           notes,
		CreatedAt:       time.Now().Format(time.RFC3339),
	}

	_, err := m.DB.Model(logs).Insert()
	if err != nil {
		return nil, err
	}

	return &model.Log{
		ID:              logs.ID,
		UserID:          logs.UserID,
		LogDate:         logs.LogDate,
		EnergyLevel:     int32(logs.EnergyLevel),
		Mood:            int32(logs.Mood),
		Libido:          int32(logs.Libido),
		SleepHours:      logs.SleepHours,
		ExerciseMinutes: logs.ExerciseMinutes,
		StressLevel:     int32(logs.StressLevel),
		Notes:           &logs.Notes,
		CreatedAt:       logs.CreatedAt,
	}, nil
}
