-- +goose Up
ALTER TABLE events
ALTER COLUMN event_cal_ids SET DEFAULT '{}'::TEXT[];

-- +goose Down
ALTER TABLE events
ALTER COLUMN event_cal_ids DROP DEFAULT;
