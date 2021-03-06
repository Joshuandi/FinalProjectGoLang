package main

import (
	//"FinalProjectGoLang/auth"
	"FinalProjectGoLang/config"
	"FinalProjectGoLang/handler"
	"FinalProjectGoLang/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
	if config.Err != nil {
		panic(config.Err)
	}
	config.Err = config.Db.Ping()
	if config.Err != nil {
		panic(config.Err)
	}
	fmt.Println("Successfully Connect to Database")

	r := mux.NewRouter()

	//user
	userHandler := handler.NewUserhandler(config.Db)
	r.HandleFunc("/users/register", userHandler.UserHandler)
	r.HandleFunc("/users/login", userHandler.UserHandler)
	r.HandleFunc("/users", userHandler.UserHandler)
	r.HandleFunc("/users/{id}", userHandler.UserHandler)
	//photo
	photoHandler := handler.NewPhotohandler(config.Db, &model.Photo{})
	r.HandleFunc("/photos", photoHandler.PhotoHandler)
	r.HandleFunc("/photos/{id}", photoHandler.PhotoHandler)
	//comment
	commentHandler := handler.NewCommenthandler(config.Db, &model.Comment{})
	r.HandleFunc("/comments", commentHandler.CommentHandler)
	r.HandleFunc("/comments/{id}", commentHandler.CommentHandler)
	//social media
	socialMediaHandler := handler.NewSocialMediahandler(config.Db, &model.SocialMedia{})
	r.HandleFunc("/socialmedias", socialMediaHandler.SocialMediaHandler)
	r.HandleFunc("/socialmedias/{id}", socialMediaHandler.SocialMediaHandler)

	//auth := auth.NewAuthMiddleware(&cfg)
	//r.Use(auth.AuthLoginValidation)
	//r.Use(auth.AuthTokenValidation)

	fmt.Println("Now Loading on Port", cfg.PORT)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8088",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
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
