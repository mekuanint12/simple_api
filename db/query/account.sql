-- name: CreateAccount :one
INSERT INTO accounts (
  user_name, 
  first_name,
  last_name,
  password,
  email
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE user_name = $1 AND password = $2
LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :exec
UPDATE accounts 
SET user_name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;
