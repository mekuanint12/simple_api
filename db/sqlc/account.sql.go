// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  user_name, 
  first_name,
  last_name,
  password,
  email
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, user_name, first_name, last_name, password, email, created_on
`

type CreateAccountParams struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.UserName,
		arg.FirstName,
		arg.LastName,
		arg.Password,
		arg.Email,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.Email,
		&i.CreatedOn,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, user_name, first_name, last_name, password, email, created_on FROM accounts
WHERE user_name = $1 AND password = $2
LIMIT 1
`

type GetAccountParams struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (q *Queries) GetAccount(ctx context.Context, arg GetAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, arg.UserName, arg.Password)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.Email,
		&i.CreatedOn,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, user_name, first_name, last_name, password, email, created_on FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserName,
			&i.FirstName,
			&i.LastName,
			&i.Password,
			&i.Email,
			&i.CreatedOn,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE accounts 
SET user_name = $2
WHERE id = $1
RETURNING id, user_name, first_name, last_name, password, email, created_on
`

type UpdateAccountParams struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.db.ExecContext(ctx, updateAccount, arg.ID, arg.UserName)
	return err
}
