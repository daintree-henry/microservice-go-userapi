package users

const(
	queryCreateUser = "INSERT INTO users(id,password,name,email,phone_number,status,created_at,modified_at) VALUES(?,?,?,?,?,?,?,?);"
	queryGetUserById = "SELECT id,name,email,phone_number,status,created_at,modified_at FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET email=?, phone_number=?, status=?, modified_at=?;"
	queryDeleteUser = "DELETE FROM users where id=?;"
	queryFindIdByPhoneNumber = "SELECT id FROM users WHERE phone_number=?;"
	queryValidUserCheck = "SELECT id,name,email,phone_number,status,created_at,modified_at FROM users WHERE id=? AND password=?"
)

func (user *User) CreateUser() errors.UtilErr{
	stmt, err := userdb.Client.Prepare(queryCreateUser)
	if err != nil{
		logger.Error("fail to prepare userdb statement", err)
		return errors.NewRestError("fail to prepare userdb statement", http.StatusInternalServerError, err.Error())
	}

	result, err := stmt.Exec(user.Id, user.Password, user.Name, user.Email, user.PhoneNumber, user.Status, user.CreatedAt, user.ModifiedAt)
	if err != nil{
		logger.Error("fail to execute queryCreateUser")
		return errors.NewRestError("fail to execute queryCreateUser", http.StatusInternalServerError, err.Error())
	}

	userKey, err := result.LastInsertId()
	if err != nil{
		logger.Error("fail to execute queryCreateUser")
		return errors.NewRestError("fail to execute queryCreateUser", http.StatusInternalServerError, err.Error())
	}
	user.PrimaryKey = userKey
	return nil
}