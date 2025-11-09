-- +goose Up
-- +goose StatementBegin
CREATE TABLE lab_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    testosterone_level FLOAT NOT NULL,
    unit VARCHAR(10) DEFAULT 'ng/dL',
    test_date DATE DEFAULT CURRENT_DATE,
    source VARCHAR(50) DEFAULT 'Manual Entry',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lab_results;
-- +goose StatementEnd
