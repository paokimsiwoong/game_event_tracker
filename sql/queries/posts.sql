-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, name, tag, tag_text, posted_at, starts_at, ends_at, body, post_url, site_id)
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
    $8,
    $9
)
RETURNING *;

-- name: CreatePostWithNull :one
INSERT INTO posts (id, created_at, updated_at, name, tag, tag_text, posted_at, body, post_url, site_id)
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
    $7
)
RETURNING *;

-- name: GetPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: GetPostsAndSites :many
SELECT posts.*, sites.name AS site_name, sites.url AS site_url FROM posts
INNER JOIN sites
ON posts.site_id = sites.id
ORDER BY posts.posted_at DESC, posts.starts_at DESC;

-- name: GetPostByID :one
SELECT * FROM posts
WHERE id = $1;

-- name: GetPostsByName :many
SELECT * FROM posts
WHERE name = $1
ORDER BY created_at;

-- name: GetPostsByNameAndPostedAtAndSiteID :many
SELECT * FROM posts
WHERE name = $1 AND posted_at = $2 AND site_id = $3
ORDER BY created_at;

-- name: GetPostsBySiteID :many
SELECT * FROM posts
WHERE site_id = $1
ORDER BY created_at;

-- name: GetPostsOnGoing :many
SELECT * FROM posts
WHERE starts_at <= NOW() AND (ends_at IS NULL OR ends_at >= NOW())
ORDER BY created_at;

-- name: GetPostsOnGoingAndSites :many
SELECT posts.*, sites.name AS site_name, sites.url AS site_url FROM posts
INNER JOIN sites
ON posts.site_id = sites.id
WHERE posts.starts_at <= NOW() AND (posts.ends_at IS NULL OR posts.ends_at >= NOW())
ORDER BY posts.posted_at DESC, posts.starts_at DESC;

-- name: GetPostsWithinGivenPeriod :many
SELECT posts.*, sites.name AS site_name, sites.url AS site_url FROM posts
INNER JOIN sites
ON posts.site_id = sites.id
WHERE posts.ends_at IS NULL OR posts.ends_at >= $1
ORDER BY posts.posted_at DESC, posts.starts_at DESC;

-- name: SetPostDates :exec
UPDATE posts
SET updated_at = NOW(), posted_at = $1, starts_at = $2, ends_at = $3
WHERE id = $4;

-- name: DeletePostByID :exec
DELETE FROM posts
WHERE id = $1;

-- name: DeletePostsBySiteID :exec
DELETE FROM posts
WHERE site_id = $1;

-- name: DeletePostsBySiteName :exec
DELETE FROM posts
USING sites
WHERE posts.site_id = sites.id AND sites.name = $1;

-- name: DeletePostsBySiteUrl :exec
DELETE FROM posts
WHERE EXISTS (
    SELECT 1 FROM sites
    WHERE sites.id = posts.site_id AND sites.url = $1
);

-- name: DeleteOldPosts :exec
DELETE FROM posts
WHERE ends_at < NOW();

-- name: ResetPosts :exec
DELETE FROM posts;
