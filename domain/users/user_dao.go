package users

import (
	"fmt"
	"net/http"

	"github.com/daintree-henry/microservice-go-userapi/database/mysql/userdb"
	"github.com/daintree-henry/microservice-go-userapi/utils/errors"
	"github.com/daintree-henry/microservice-go-userapi/utils/logger"
)

const (
	queryCreateUser          = "INSERT INTO users(id,password,name,email,phone_number,status,created_at,modified_at) VALUES(?,?,?,?,?,?,?,?);"
	queryGetUserByPK         = "SELECT id,name,email,phone_number,status,created_at,modified_at FROM users WHERE id = ?;"
	queryPrimaryKeyById      = "SELECT primary_key FROM users WHERE id = ?"
	queryUpdateUserByPK      = "UPDATE users SET email=?, phone_number=?, status=?, modified_at=? WHERE primary_key = ?;"
	queryDeleteUserByPK      = "DELETE FROM users WHERE primary_key = ?;"
	queryFindIdByPhoneNumber = "SELECT id FROM users WHERE phone_number=?;"
	queryValidUserCheck      = "SELECT id,name,email,phone_number,status,created_at,modified_at FROM users WHERE id=? AND password=? AND status=?"
	queryGetUsersByStatus    = "SELECT id,name,email,phone_number,status,created_at,modified_at FROM users WHERE status = ?"
)

func (user *User) CreateUser() errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryCreateUser)
	if err != nil {
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Id, user.Password, user.Name, user.Email, user.PhoneNumber, user.Status, user.CreatedAt, user.ModifiedAt)
	if err != nil {
		logger.Error("fail to execute queryCreateUser", err)
		return errors.NewRestError("fail to execute queryCreateUser", http.StatusInternalServerError, err.Error())
	}

	userKey, err := result.LastInsertId()
	if err != nil {
		logger.Error("fail to execute queryCreateUser", err)
		return errors.NewRestError("fail to execute queryCreateUser", http.StatusInternalServerError, err.Error())
	}
	user.PrimaryKey = userKey
	logger.Info(fmt.Sprintf("user made successfully : ", userKey))
	return nil
}

func (user *User) GetUserById() errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryGetUserById)
	if err != nil {
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	err := stmt.QueryRow(user.PrimaryKey).Scan(&user.Id, &user.Name, &user.Email, &user.PhoneNumber, &user.Status, &user.CreatedAt, &user.ModifiedAt)
	if err != nil {
		logger.Error("fail to execute queryGetUserById", err)
		return errors.NewRestError("fail to execute queryGetUserById", http.StatusInternalServerError, err.Error())
	}

	logger.Info("getting user successfully")
	return nil
}

func (user *User) PrimaryKeyById() errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryPrimaryKeyById)
	if err != nil {
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	err := stmt.QueryRow(user.PrimaryKey).Scan(&user.Id)
	if err != nil {
		logger.Error("fail to retrieve user", err)
		return errors.NewRestError("fail to retrieve user", http.StatusInternalServerError, err.Error())
	}

	logger.Info("getting user Id successfully")
	return nil
}

func (user *User) UpdateUserByPK() errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryUpdateUserByPK)
	if err != nil {
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.PhoneNumber, user.Status, user.ModifiedAt, user.PrimaryKey)
	if err != nil {
		logger.Error("fail to execute queryUpdateUserByPK", err)
		return errors.NewRestError("fail to execute queryUpdateUserByPK", http.StatusInternalServerError, err.Error())
	}

	logger.Error("update user successfully", err)
	return nil
}

func (user *User) DeleteUserByPK() errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryDeleteUserByPK)
	if err != nil {
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.PrimaryKey); err != nil {
		logger.Error("fail to execute queryDeleteUserByPK", err)
		return errors.NewRestError("fail to execute queryDeleteUserByPK", http.StatusInternalServerError, err.Error())
	}
	logger.Error("delete user successfully", err)
	return nil
}

func (user *User) FindIdByPhoneNumber() errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryFindIdByPhoneNumber)
	if err != nil {
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	if err := stmt.QueryRow(user.PhoneNumber).Scan(&user); err != nil {
		logger.Error("fail to execute queryFindIdByPhoneNumber", err)
		return errors.NewRestError("fail to execute queryFindIdByPhoneNumber", http.StatusInternalServerError, err.Error())
	}
	logger.Error("get user successfully", err)
	return nil
}

func (user *User) queryValidUserCheck() errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryValidUserCheck)
	if err != nil {
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.PhoneNumber, user.Password)
	if err := result.Scan(); err != nil {
		logger.Error("fail to execute queryValidUserCheck", err)
		return errors.NewRestError("fail to execute queryValidUserCheck", http.StatusUnauthorized, err.Error())
	}
	logger.Info("valid id and password")
	return nil
}

func (user *User) GetUsersByStatus(status string) ([]User, errors.UtilErr) {
	stmt, err := userdb.Client.Prepare(queryGetUsersByStatus)
	if err != nil {
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("fail to execute queryGetUsersByStatus", err)
		return errors.NewRestError("fail to execute queryGetUsersByStatus", http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	results = make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.PhoneNumber, &user.Status, &user.CreatedAt, &user.ModifiedAt); err != nil {
			logger.Error("fail to scanning user rows", err)
			return errors.NewRestError("fail to scanning user rows", http.StatusInternalServerError, err.Error())
		}
		results.append(results, user)
	}
	if len(results) == nil {
		logger.Error(fmt.Sprintf("there is no ", status, " user"), nil)
		return errors.NewRestError(fmt.Sprintf("there is no ", status, " user"), http.StatusInternalServerError, err.Error())
	}
	return results, nil
}
