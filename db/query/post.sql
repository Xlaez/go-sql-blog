-- name: CreatePost :one
INSERT INTO post (
  title,
  description,
  content,
  user_name
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetPost :one
SELECT * FROM post
WHERE id = $1 LIMIT 1;

-- name: GetPostForUpdate :one
SELECT * FROM post
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- name: ListPosts :many
SELECT * FROM post
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListPostsForUser :many
SELECT * FROM "user"
JOIN post ON post.user_name $1
LIMIT $2
OFFSET $3;

-- name: UpdatePost :one
UPDATE post 
SET content = $2
WHERE id = $1
RETURNING *; 

-- name: DeletePost :exec
DELETE FROM post WHERE id = $1;