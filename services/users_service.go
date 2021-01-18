package services

import (
	"net/http"

	"github.com/daintree-henry/microservice-go-userapi/domain/users"
	"github.com/daintree-henry/microservice-go-userapi/utils/utils_crypto"
	"github.com/daintree-henry/microservice-go-userapi/utils/utils_date"
	"github.com/daintree-henry/microservice-go-userapi/utils/utils_errors"
)

//인터페이스에 구조체 주입
var UsersService usersServiceInterface = &usersService{}

//인터페이스를 구현한 구조체
type usersService struct{}

type usersServiceInterface interface {
	SvcCreateUser(users.User) (*users.User, utils_errors.UtilErr)
	SvcGetUserById(string) (*users.User, utils_errors.UtilErr)
	SvcUpdateUser(bool, users.User) (*users.User, utils_errors.UtilErr)
	SvcDeleteUserById(string) utils_errors.UtilErr
	SvcLoginUserByIdAndPw(users.User) (*users.User, utils_errors.UtilErr)
	SvcFindIdByPhoneNumber(users.User) (*users.User, utils_errors.UtilErr)
	SvcGetUsersByStatus(string) (*[]users.User, utils_errors.UtilErr)
}

func (s *usersService) SvcCreateUser(user users.User) (*users.User, utils_errors.UtilErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	nowTime := utils_date.GetNowDateString()
	user.CreatedAt = nowTime
	user.ModifiedAt = nowTime
	cryptoPassword, err := utils_crypto.HashPassword(user.Password)
	if err != nil {
		return nil, utils_errors.NewRestError("Error in hashing password", http.StatusInternalServerError, err.Error())
	}
	user.Password = cryptoPassword

	if err := user.DaoCreateUser(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) SvcGetUserById(userid string) (*users.User, utils_errors.UtilErr) {
	dao := &users.User{Id: userid}
	if err := dao.DaoPrimaryKeyById(); err != nil {
		return nil, err
	}
	if err := dao.DaoGetUserByPK(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *usersService) SvcUpdateUser(isPartial bool, user users.User) (*users.User, utils_errors.UtilErr) {
	dao := &users.User{Id: user.Id}
	if err := dao.DaoPrimaryKeyById(); err != nil {
		return nil, err
	}
	if err := dao.DaoGetUserByPK(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.Password != "" {
			dao.Password = user.Password
		}
		if user.Name != "" {
			dao.Name = user.Name
		}
		if user.Email != "" {
			dao.Email = user.Email
		}
		if user.PhoneNumber != "" {
			dao.PhoneNumber = user.PhoneNumber
		}
		if user.Status != "" {
			dao.Status = user.Status
		}
	} else {
		dao.Password = user.Password
		dao.Name = user.Name
		dao.Email = user.Email
		dao.PhoneNumber = user.PhoneNumber
		dao.Status = user.Status
	}
	dao.ModifiedAt = utils_date.GetNowDateString()

	if err := dao.DaoUpdateUserByPK(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *usersService) SvcDeleteUserById(userid string) utils_errors.UtilErr {
	dao := &users.User{Id: userid}
	if err := dao.DaoPrimaryKeyById(); err != nil {
		return err
	}
	if err := dao.DaoDeleteUserByPK(); err != nil {
		return err
	}
	return nil
}

func (s *usersService) SvcLoginUserByIdAndPw(user users.User) (*users.User, utils_errors.UtilErr) {
	dao := &users.User{
		Id:       user.Id,
		Password: user.Password,
	}
	if err := dao.DaoValidUserCheck(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *usersService) SvcFindIdByPhoneNumber(user users.User) (*users.User, utils_errors.UtilErr) {
	dao := &users.User{PhoneNumber: user.PhoneNumber}
	if err := dao.DaoFindIdByPhoneNumber(); err != nil {
		return nil, err
	}

	if dao.Id == "" {
		return nil, utils_errors.NewRestError("No matched user", http.StatusNotFound, "")
	}
	return dao, nil
}

func (s *usersService) SvcGetUsersByStatus(status string) (*[]users.User, utils_errors.UtilErr) {
	dao := &users.User{}
	return dao.DaoGetUsersByStatus(status)
}
