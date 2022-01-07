package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iniyusril/template/helper"
	"github.com/iniyusril/template/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into public.user (id,username,password) values ($1,$2,$3) RETURNING id"

	var id string
	if err := tx.QueryRowContext(ctx, SQL, user.ID, user.Username, user.Password).Scan(&id); err != nil {
		helper.PanicIfError(err)

	}

	user.ID = id
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update user set username = $1, password $2 where id = $3"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.ID)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from user where id = $1"
	_, err := tx.ExecContext(ctx, SQL, user.ID)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (domain.User, error) {
	SQL := "select id, username, password from user where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		err := errors.New("user is not found")
		helper.PanicIfError(err)
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, username,password from user"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
