-- name: CreateEvent :one
INSERT INTO events (id, created_at, updated_at, name, tag, tag_text, posted_at, starts_at, ends_at, body, site_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: CreateEventWithNull :one
INSERT INTO events (id, created_at, updated_at, name, tag, tag_text, body, site_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetEvents :many
SELECT * FROM events
ORDER BY created_at;

-- name: GetEventByID :one
SELECT * FROM events
WHERE id = $1;

-- name: GetEventByName :one
SELECT * FROM events
WHERE name = $1;

-- name: GetEventByNameAndPostedAtAndSiteID :one
SELECT * FROM events
WHERE name = $1 AND posted_at = $2 AND site_id = $3;

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
SET updated_at = NOW(), posted_at = $1, starts_at = $2, ends_at = $3
WHERE id = $4;

-- name: DeleteEventByID :exec
DELETE FROM events
WHERE id = $1;

-- name: DeleteEventsBySiteID :exec
DELETE FROM events
WHERE site_id = $1;

-- name: DeleteOldEvents :exec
DELETE FROM events
WHERE ends_at < NOW();
