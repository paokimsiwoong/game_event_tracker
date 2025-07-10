-- name: CreateEvent :one
INSERT INTO events (id, created_at, updated_at, name, tag, body, site_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetEvents :many
SELECT * FROM events
ORDER BY created_at;

-- name: GetEventByID :one
SELECT * FROM events
WHERE id = $1;

-- name: GetEventsBySiteID :many
SELECT * FROM events
WHERE site_id = $1
ORDER BY created_at;

-- name: GetEventsOnGoing :many
SELECT * FROM events
WHERE starts_at <= NOW() and ends_at >= NOW()
ORDER BY created_at;

-- name: SetEventDates :exec
UPDATE events
SET updated_at = NOW(), starts_at = $1, ends_at = $2
WHERE id = $3;

-- name: DeleteEventByID :exec
DELETE FROM events
WHERE id = $1;

-- name: DeleteEventsBySiteID :exec
DELETE FROM events
WHERE site_id = $1;

-- name: DeleteOldEvents :exec
DELETE FROM events
WHERE ends_at < NOW();
