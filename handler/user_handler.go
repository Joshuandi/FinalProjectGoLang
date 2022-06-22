package handler

import (
	"FinalProjectGoLang/config"
	user "FinalProjectGoLang/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type UserHandlerInterface interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}
type UserHandler struct {
	db *sql.DB
}

func NewUserhandler(db *sql.DB) UserHandlerInterface {
	return &UserHandler{db: db}
}

func (u *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	switch r.Method {
	case http.MethodGet:
		u.UserGetAll(w, r)
	case http.MethodPost:
		u.UserPostRegister(w, r)
	case http.MethodPut:
		u.UserUpdate(w, r, id)
	case http.MethodDelete:
		u.UserDelete(w, r, id)
	}
}

func (u *UserHandler) UserGetAll(w http.ResponseWriter, r *http.Request) {
	var result = []user.User{}
	var userss = user.User{}
	sqlGet := "Select u_id, u_username, u_email, u_age from users;"
	rows, err := config.Db.Query(sqlGet)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&userss.User_id,
			&userss.Username,
			&userss.Email,
			&userss.Age,
		); err != nil {
			fmt.Println("No Data", err)
		}
		result = append(result, userss)
	}
	jsonData, _ := json.Marshal(&result)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
func (u *UserHandler) UserPostRegister(w http.ResponseWriter, r *http.Request) {
	var users = user.User{}
	// if users.Email == "" || !strings.Contains(users.Email, "@gmail.com") ||
	// 	users.Username == "" || users.Password == "" || len(users.Password) < 6 ||
	// 	users.Age == 0 || users.Age <= 8 {
	// 	errors.New("Data harus di isi semua")
	// } else {
	json.NewDecoder(r.Body).Decode(&users)
	password := []byte(users.Password)
	sqlSt := `insert into users (u_username, u_email, u_pass, u_age,u_created_date, u_updated_date)
		values ($1, $2, $3, $4, $5, $6)
		returning u_id;`

	HashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	err := config.Db.QueryRow(sqlSt,
		users.Username,
		users.Email,
		string(HashPassword),
		users.Age,
		time.Now(),
		time.Now(),
	).Scan(&users.User_id)
	if err != nil {
		panic(err)
	}
	Register_respone := user.RegisterRespone{
		R_user_id:  users.User_id,
		R_email:    users.Email,
		R_username: users.Username,
		R_age:      users.Age,
	}
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

	fmt.Println(users)
	fmt.Println(Register_respone)
	return
	//}
}

func (u *UserHandler) UserPostLogin(w http.ResponseWriter, r *http.Request)         {}
func (u *UserHandler) UserUpdate(w http.ResponseWriter, r *http.Request, id string) {}
func (u *UserHandler) UserDelete(w http.ResponseWriter, r *http.Request, id string) {
	sqlDelete := `DELETE from users WHERE u_id = $1`
	if index, err := strconv.Atoi(id); err == nil {
		res, err := config.Db.Exec(sqlDelete, index)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(fmt.Sprint("Deleted Data", count)))
		return
	}
}
