package main

import (
	"FinalProjectGoLang/config"
	"database/sql"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

var cfg config.Config

func main() {
	_ = cleanenv.ReadConfig(".env", &cfg)
	config.Db, config.Err = sql.Open("postgres", ConnectDbPsql(
		cfg.Db_Host,
		cfg.Db_Dbname,
		cfg.Db_User,
		cfg.Db_Password,
		cfg.Db_Port,
	))
	defer config.Db.Close()

	config.Err = config.Db.Ping()
	if config.Err != nil {
		panic(config.Err)
	}
	fmt.Println("Successfully Connect to Database")
}

func ConnectDbPsql(host, user, password, dbname string, port int) string {
	_ = cleanenv.ReadConfig(".env", &cfg)
	psqlInfo := fmt.Sprintf("host= %s port= %d user= %s "+
		" password= %s dbname= %s sslmode=disable",
		cfg.Db_Host,
		cfg.Db_Port,
		cfg.Db_User,
		cfg.Db_Password,
		cfg.Db_Dbname)
	return psqlInfo
}
