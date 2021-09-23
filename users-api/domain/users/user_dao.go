package users

import (
	"fmt"
	"github.com/sijanstha/datasources/mysql/users_db"
	"github.com/sijanstha/utils/errors"
)

const (
	queryInsertUser     = "insert into users(first_name, last_name, email, date_created) values (?, ?, ?, ?);"
	queryGetUserById    = "select id, first_name, last_name, email, date_created from users where id=?;"
	queryGetUserByEmail = "select id, first_name, last_name, email, date_created from users where email=?;"
	queryUpdateUser     = "update users set first_name=?, last_name=?, email=? where id=?;"
	queryDeleteUser     = "delete from users where id=?"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
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

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
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
	err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
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
