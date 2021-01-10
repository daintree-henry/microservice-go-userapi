package users

import (
	"fmt"
	"net/http"

	"github.com/daintree-henry/microservice-go-userapi/database/mysql/userdb"
	"github.com/daintree-henry/microservice-go-userapi/domain/users"
	"github.com/daintree-henry/microservice-go-userapi/utils/utils_errors"
	"github.com/daintree-henry/microservice-go-userapi/utils/utils_logger"
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

func (user *User) DaoCreateUser() utils_errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryCreateUser)
	if err != nil {
		utils_logger.Error("fail to prepare userdb statement", err)
		return utils_errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Id, user.Password, user.Name, user.Email, user.PhoneNumber, user.Status, user.CreatedAt, user.ModifiedAt)
	if err != nil {
		utils_logger.Error("fail to execute queryCreateUser", err)
		return utils_errors.NewRestError("fail to execute queryCreateUser", http.StatusInternalServerError, err.Error())
	}

	userKey, err := result.LastInsertId()
	if err != nil {
		utils_logger.Error("fail to execute queryCreateUser", err)
		return utils_errors.NewRestError("fail to execute queryCreateUser", http.StatusInternalServerError, err.Error())
	}
	user.PrimaryKey = userKey
	utils_logger.Info(fmt.Sprintf("user made successfully : ", userKey))
	return nil
}

func (user *User) DaoGetUserByPK() utils_errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryGetUserByPK)
	if err != nil {
		utils_logger.Error("fail to prepare userdb statement", err)
		return utils_errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	err := stmt.QueryRow(user.PrimaryKey).Scan(&user.Id, &user.Name, &user.Email, &user.PhoneNumber, &user.Status, &user.CreatedAt, &user.ModifiedAt)
	if err != nil {
		utils_logger.Error("fail to execute queryGetUserById", err)
		return utils_errors.NewRestError("fail to execute queryGetUserById", http.StatusInternalServerError, err.Error())
	}

	utils_logger.Info("getting user successfully")
	return nil
}

func (user *User) DaoPrimaryKeyById() utils_errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryPrimaryKeyById)
	if err != nil {
		utils_logger.Error("fail to prepare userdb statement", err)
		return utils_errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	err := stmt.QueryRow(user.PrimaryKey).Scan(&user.Id)
	if err != nil {
		utils_logger.Error("fail to retrieve user", err)
		return utils_errors.NewRestError("fail to retrieve user", http.StatusInternalServerError, err.Error())
	}

	utils_logger.Info("getting user PK successfully")
	return nil
}

func (user *User) DaoUpdateUserByPK() utils_errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryUpdateUserByPK)
	if err != nil {
		utils_logger.Error("fail to prepare userdb statement", err)
		return utils_errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.PhoneNumber, user.Status, user.ModifiedAt, user.PrimaryKey)
	if err != nil {
		utils_logger.Error("fail to execute queryUpdateUserByPK", err)
		return utils_errors.NewRestError("fail to execute queryUpdateUserByPK", http.StatusInternalServerError, err.Error())
	}

	utils_logger.Error("update user successfully", err)
	return nil
}

func (user *User) DaoDeleteUserByPK() utils_errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryDeleteUserByPK)
	if err != nil {
		utils_logger.Error("fail to prepare userdb statement", err)
		return utils_errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.PrimaryKey); err != nil {
		utils_logger.Error("fail to execute queryDeleteUserByPK", err)
		return utils_errors.NewRestError("fail to execute queryDeleteUserByPK", http.StatusInternalServerError, err.Error())
	}
	utils_logger.Error("delete user successfully", err)
	return nil
}

func (user *User) DaoFindIdByPhoneNumber() utils_errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryFindIdByPhoneNumber)
	if err != nil {
		utils_logger.Error("fail to prepare userdb statement", err)
		return utils_errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	if err := stmt.QueryRow(user.PhoneNumber).Scan(&user); err != nil {
		utils_logger.Error("fail to execute queryFindIdByPhoneNumber", err)
		return utils_errors.NewRestError("fail to execute queryFindIdByPhoneNumber", http.StatusInternalServerError, err.Error())
	}
	utils_logger.Error("get user successfully", err)
	return nil
}

func (user *User) DaoValidUserCheck() utils_errors.UtilErr {
	stmt, err := userdb.Client.Prepare(queryValidUserCheck)
	if err != nil {
		utils_logger.Error("fail to prepare userdb statement", err)
		return utils_errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	"SELECT id,name,email,phone_number,status,created_at,modified_at FROM users WHERE id=? AND password=? AND status=?"

	result := stmt.QueryRow(user.PhoneNumber, user.Password, users.StatusActive)
	if err := result.Scan(&user.Id, &user.Name, &user.Email, &user.PhoneNumber, &user.Status, &user.CreatedAt, &user.ModifiedAt); err != nil {
		utils_logger.Error("fail to execute queryValidUserCheck", err)
		return utils_errors.NewRestError("fail to execute queryValidUserCheck", http.StatusUnauthorized, err.Error())
	}
	utils_logger.Info("valid id and password")
	return nil
}

func (user *User) DaoGetUsersByStatus(status string) ([]User, utils_errors.UtilErr) {
	stmt, err := userdb.Client.Prepare(queryGetUsersByStatus)
	if err != nil {
		utils_logger.Error("fail to prepare userdb statement", err)
		return utils_errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		utils_logger.Error("fail to execute queryGetUsersByStatus", err)
		return utils_errors.NewRestError("fail to execute queryGetUsersByStatus", http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	results = make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.PhoneNumber, &user.Status, &user.CreatedAt, &user.ModifiedAt); err != nil {
			utils_logger.Error("fail to scanning user rows", err)
			return utils_errors.NewRestError("fail to scanning user rows", http.StatusInternalServerError, err.Error())
		}
		results.append(results, user)
	}
	if len(results) == nil {
		utils_logger.Error(fmt.Sprintf("there is no ", status, " user"), nil)
		return utils_errors.NewRestError(fmt.Sprintf("there is no ", status, " user"), http.StatusInternalServerError, err.Error())
	}
	return results, nil
}
