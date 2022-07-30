package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type CreateUser struct {
	*Queries
	db *sql.DB
}

func NewUser(db *sql.DB) *CreateUser {
	return &CreateUser{
		db:      db,
		Queries: New(db),
	}
}

func (create *CreateUser) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := create.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type SavedResult struct {
	ID        Account `json:"id"`
	UserName  Account `json:"user_name"`
	FirstName Account `json:"first_name"`
	LastName  Account `json:"last_name"`
	Password  Account `json:"password"`
	Email     Account `json:"email"`
	CreatedOn Account `json:"created_on"`
}

func (cUser *CreateUser) SaveTx(ctx context.Context, arg CreateAccountParams) (SavedResult, error) {
	var result SavedResult

	err := cUser.execTx(ctx, func(q *Queries) error {
		var err error
		result.ID, err = q.CreateAccount(ctx, CreateAccountParams{
			UserName:  arg.UserName,
			FirstName: arg.FirstName,
			LastName:  arg.LastName,
			Password:  arg.Password,
			Email:     arg.Email,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
