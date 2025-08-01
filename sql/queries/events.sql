-- name: CreateEvent :one
INSERT INTO events (id, created_at, updated_at, tag, tag_text, starts_at, ends_at, names, posted_ats, post_urls, post_ids, site_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    ARRAY[$5],
    ARRAY[$6],
    ARRAY[$7],
    ARRAY[$8],
    $9
)
ON CONFLICT (tag, starts_at, ends_at)
DO UPDATE SET names = array_append(names, $5), 
posted_ats = array_append(posted_ats, $6),
post_urls = array_append(post_urls, $7),
post_ids = array_append(post_ids, $8),
updated_at = NOW()
RETURNING *;

-- name: GetEventByID :one
SELECT * FROM events
WHERE id = $1;

-- name: GetEvents :many
SELECT * FROM events
ORDER BY created_at;

-- name: GetEventsByTag :many
SELECT * FROM events
WHERE tag = $1
ORDER BY created_at;

-- name: GetEventsByTagText :many
SELECT * FROM events
WHERE tag_text = $1
ORDER BY created_at;

-- name: GetEventsBySiteID :many
SELECT * FROM events
WHERE site_id = $1
ORDER BY created_at;

-- name: GetEventsOnGoing :many
SELECT * FROM events
WHERE starts_at <= NOW() AND (ends_at IS NULL OR ends_at >= NOW())
ORDER BY created_at;

-- name: GetOldEvents :many
SELECT * FROM events
WHERE ends_at < NOW()
ORDER BY created_at;

-- name: SetEventCalID :exec
UPDATE events
SET updated_at = NOW(), event_cal_id = $1
WHERE tag = $2 AND starts_at = $3 AND ends_at = $4;

-- name: SetEventCalIDByID :exec
UPDATE events
SET updated_at = NOW(), event_cal_id = $1
WHERE id = $2;

-- name: DeleteEventByID :exec
DELETE FROM events
WHERE id = $1;

-- name: DeleteEventsBySiteID :exec
DELETE FROM events
WHERE site_id = $1;

-- name: DeleteEventBySiteName :exec
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