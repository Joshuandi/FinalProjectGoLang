package config

import "database/sql"

var (
	Db  *sql.DB
	Err error
)

type Config struct {
	Lusername      string `env:"Lusername"`
	Lpassword      string `env:"Lpassword"`
	GoogleClientId string `env:"Google_client_id"`
	Db_Host        string `env:"Host"`
	Db_Port        int    `env:"Port"`
	Db_User        string `env:"User"`
	Db_Password    string `env:"Password"`
	Db_Dbname      string `env:"Dbname"`
	PORT           string `env:PORT`
}
