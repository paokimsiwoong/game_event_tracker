-- returning 으로 생성한 유저를 바로 반환하고 있음 (위에 :one으로 생성한 유저 하나만 반환하도록 함)
-- name: CreateSite :one
INSERT INTO sites (id, created_at, updated_at, name, url)
VALUES (
    gen_random_uuid(), 
    NOW(),
    NOW(),
    $1,
    $2
)   
RETURNING *; 

-- name: GetSiteByName :one
SELECT * FROM sites 
WHERE name = $1; 

-- name: GetSiteByURL :one
SELECT * FROM sites
WHERE url = $1;

-- name: GetSites :many
SELECT *
FROM sites
ORDER BY sites.updated_at;

-- name: MarkSiteFetched :exec
UPDATE sites
SET updated_at = NOW(), last_fetched_at = NOW()
WHERE id = $1;

-- name: DeleteSiteByName :exec
DELETE FROM sites
WHERE name = $1;

-- name: DeleteSiteByURL :exec
DELETE FROM sites
WHERE url = $1;

-- name: ResetSites :exec
DELETE FROM sites;