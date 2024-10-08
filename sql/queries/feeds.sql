-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsByUserID :many
SELECT * FROM feeds WHERE user_id = $1;