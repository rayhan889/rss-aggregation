-- name: CreateUser :one
INSERT INTO users (id, email, name, password) 
VALUES ($1, $2, $3, $4) 
RETURNING *;