package users

import (
	"fmt"

	"github.com/sijanstha/common-utils/src/logger"
	"github.com/sijanstha/common-utils/src/utils/errors"
	"github.com/sijanstha/datasources/mysql/users_db"
)

const (
	queryInsertUser = "insert into users(first_name, last_name, email, date_created, status, password) values (?, ?, ?, ?, ?, ?);"
	queryUpdateUser = "update users set first_name=?, last_name=?, email=? where id=?;"
	queryDeleteUser = "delete from users where id=?;"
)

func (user *User) Find(filter UserFilter) *errors.RestErr {

	query, args := prepareSelectQuery(filter)

	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError(errors.NewError("database error").Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(args...)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found.", user.Id))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (user *User) FindAll(filter UserFilter) ([]User, *errors.RestErr) {

	query, args := prepareSelectQuery(filter)
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}

	return results, nil
}

func prepareSelectQuery(filter UserFilter) (string, []interface{}) {
	args := make([]interface{}, 0)

	query := "select id, first_name, last_name, email, date_created, status from users where 1=1 "
	if filter.Id != 0 {
		query += "and id = ? "
		args = append(args, filter.Id)
	}

	if filter.Email != "" {
		query += "and email = ? "
		args = append(args, filter.Email)
	}

	if filter.Password != "" {
		query += "and password = ? "
		args = append(args, filter.Password)
	}

	if filter.Status != "" {
		query += "and status = ? "
		args = append(args, filter.Status)
	}

	logger.Info(query)
	return query, args
}
