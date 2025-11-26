package dbmodels

type Log struct {
	tableName struct{} `pg:"logs"` // table name

	ID              string  `pg:"id,pk"`
	UserID          string  `pg:"user_id"`
	LogDate         string  `pg:"log_date"`
	EnergyLevel     int     `pg:"energy_level"`
	Mood            int     `pg:"mood"`
	Libido          int     `pg:"libido"`
	SleepHours      float64 `pg:"sleep_hours"`
	ExerciseMinutes float64 `pg:"exercise_minutes"`
	StressLevel     int     `pg:"stress_level"`
	Notes           string  `pg:"notes"`
	CreatedAt       string  `pg:"created_at"`
}
