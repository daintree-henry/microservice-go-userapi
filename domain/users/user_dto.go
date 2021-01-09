package users

import (
	"strings"

	"github.com/daintree-henry/microservice-go-userapi/utils/errors"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

//User 구조체와 필드 정의
type User struct {
	PrimaryKey  int64  `json:"primary_key"`
	Id          string `json:"id"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	ModifiedAt  string `json:"modified_at"`
}

//User의 리스트인 Users 정의
type Users []User

//데이터 유효성 검사
func (user *User) Validate() errors.UtilErr {
	//marshal된 json 데이터의 공백 제거
	user.Id = strings.TrimSpace(user.Id)
	user.Password = strings.TrimSpace(user.Password)
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.PhoneNumber = strings.TrimSpace(user.PhoneNumber)

	return nil
}
