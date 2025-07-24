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

-- name: GetEventsAndSites :many
SELECT events.*, sites.name AS site_name, sites.url AS site_url FROM events
INNER JOIN sites
ON events.site_id = sites.id
ORDER BY events.posted_at DESC, events.starts_at DESC;

-- name: GetEventByID :one
SELECT * FROM events
WHERE id = $1;

-- name: GetEventsByName :many
SELECT * FROM events
WHERE name = $1
ORDER BY created_at;

-- name: GetEventsByNameAndPostedAtAndSiteID :many
SELECT * FROM events
WHERE name = $1 AND posted_at = $2 AND site_id = $3
ORDER BY created_at;

-- name: GetEventsBySiteID :many
SELECT * FROM events
WHERE site_id = $1
ORDER BY created_at;

-- name: GetEventsOnGoing :many
SELECT * FROM events
WHERE starts_at <= NOW() and ends_at >= NOW()
ORDER BY created_at;

-- name: GetEventsOnGoingAndSites :many
SELECT events.*, sites.name AS site_name, sites.url AS site_url FROM events
INNER JOIN sites
ON events.site_id = sites.id
WHERE events.starts_at <= NOW() AND (events.ends_at IS NULL OR events.ends_at >= NOW())
ORDER BY events.posted_at DESC, events.starts_at DESC;

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

-- name: DeleteEventsBySiteName :exec
DELETE FROM events
USING sites
WHERE events.site_id = sites.id AND sites.name = $1;

-- name: DeleteEventsBySiteUrl :exec
DELETE FROM events
WHERE EXISTS (
    SELECT 1 FROM sites
    WHERE sites.id = events.site_id AND sites.url = $1
);

-- name: DeleteOldEvents :exec
DELETE FROM events
WHERE ends_at < NOW();

-- name: ResetEvents :exec
DELETE FROM events;
