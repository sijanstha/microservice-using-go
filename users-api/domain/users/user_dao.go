package users

import (
	"fmt"
	"github.com/sijanstha/datasources/mysql/users_db"
	"github.com/sijanstha/logger"
	"github.com/sijanstha/utils/errors"
)

const (
	queryInsertUser             = "insert into users(first_name, last_name, email, date_created, status, password) values (?, ?, ?, ?, ?, ?);"
	queryGetUserById            = "select id, first_name, last_name, email, date_created, status from users where id=?;"
	queryGetUserByEmail         = "select id, first_name, last_name, email, date_created, status from users where email=?;"
	queryUpdateUser             = "update users set first_name=?, last_name=?, email=? where id=?;"
	queryDeleteUser             = "delete from users where id=?;"
	queryFindByStatus           = "select id, first_name, last_name, email, date_created, status from users where status=?;"
	queryFindByEmailAndPassword = "select id, first_name, last_name, email, date_created, status from users where email=? and password=? and status=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError(errors.NewError("database error").Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
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

func (user *User) FindByEmail() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByEmail)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with email %s not found", user.Email))
	}

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

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
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
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password", err)
		return errors.NewInternalServerError(errors.NewError("database error").Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return errors.NewNotFoundError("invalid user credentials")
	}

	return nil
}
