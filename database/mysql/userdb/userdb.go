package userdb

import (
	"database/sql"
	"fmt"

	"github.com/daintree-henry/microservice-go-userapi/utils/logger"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUser     = "MYSQL_USER"
	mysqlPassword = "MYSQL_PASSWORD"
	mysqlHost     = "MYSQL_HOST"
	mysqlDatabase = "MYSQL_DATABASE"
)

var (
	Client *sql.DB

	// username = os.Getenv(mysqlUser)
	// password = os.Getenv(mysqlPassword)
	// host     = os.Getenv(mysqlHost)
	// schema   = os.Getenv(mysqlDatabase)

	//localhost
	username = "root"
	password = "root"
	host     = "127.0.0.1:3306"
	schema   = "userdb"
)

func init() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	//username:password@tcp(host)/schema?charset=utf8
	var err error
	Client, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	logger.Info("Database Connected")
}
