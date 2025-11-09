-- +goose Up
-- +goose StatementBegin
CREATE TABLE logs(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    log_date DATE DEFAULT CURRENT_DATE,
    energy_level INT CHECK (energy_level BETWEEN 1 AND 5),
    mood INT CHECK (mood BETWEEN 1 AND 5),
    libido INT CHECK (libido BETWEEN 1 AND 5),
    sleep_hours FLOAT,
    exercise_minutes FLOAT,
    stress_level INT,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS logs;
-- +goose StatementEnd
