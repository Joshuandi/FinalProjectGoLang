package handler

import (
	"FinalProjectGoLang/config"
	"FinalProjectGoLang/model"
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
	db    *sql.DB
	users *model.User
}

func NewUserhandler(db *sql.DB, users *model.User) UserHandlerInterface {
	return &UserHandler{db: db, users: users}
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
	var result = []model.User{}
	sqlGet := "Select u_id, u_username, u_email, u_age from users;"
	rows, err := config.Db.Query(sqlGet)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&u.users.User_id,
			&u.users.Username,
			&u.users.Email,
			&u.users.Age,
		); err != nil {
			fmt.Println("No Data", err)
		}
		result = append(result, *u.users)
	}
	jsonData, _ := json.Marshal(&result)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
func (u *UserHandler) UserPostRegister(w http.ResponseWriter, r *http.Request) {
	// if users.Email == "" || !strings.Contains(users.Email, "@gmail.com") ||
	// 	users.Username == "" || users.Password == "" || len(users.Password) < 6 ||
	// 	users.Age == 0 || users.Age <= 8 {
	// 	errors.New("Data harus di isi semua")
	// } else {
	json.NewDecoder(r.Body).Decode(&u.users)
	password := []byte(u.users.Password)
	sqlSt := `insert into users (u_username, u_email, u_pass, u_age,u_created_date, u_updated_date)
		values ($1, $2, $3, $4, $5, $6)
		returning u_id;`

	HashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	err := config.Db.QueryRow(sqlSt,
		u.users.Username,
		u.users.Email,
		string(HashPassword),
		u.users.Age,
		time.Now(),
		time.Now(),
	).Scan(&u.users.User_id)
	if err != nil {
		panic(err)
	}
	Register_respone := model.UserRegisterRespone{
		R_user_id:  u.users.User_id,
		R_email:    u.users.Email,
		R_username: u.users.Username,
		R_age:      u.users.Age,
	}
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

	fmt.Println(u.users)
	fmt.Println(Register_respone)
	return
	//}
}

func (u *UserHandler) UserPostLogin(w http.ResponseWriter, r *http.Request) {}
func (u *UserHandler) UserUpdate(w http.ResponseWriter, r *http.Request, id string) {
	for id != "" {

		//var orders = order.Order{}
		json.NewDecoder(r.Body).Decode(&u.users)
		sqlSt := `update users set 
		u_username = $2, u_email = $3, u_updated_date = $4 where u_id = $1;`
		res, err := config.Db.Exec(sqlSt,
			id,
			u.users.Username,
			u.users.Email,
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
		count, err := res.RowsAffected()
		if err != nil {
			panic(err)
		}
		w.Write([]byte(fmt.Sprintln("Update data :", count)))
		Update_respone := model.UserUpdateRespone{
			U_user_id:    u.users.User_id,
			U_email:      u.users.Email,
			U_username:   u.users.Username,
			U_age:        u.users.Age,
			U_Updated_at: u.users.Updated_at,
		}
		jsonData, _ := json.Marshal(Update_respone)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		return
	}
}

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
		w.Write([]byte(fmt.Sprint("Message: Your account has been successfully deleted ", count)))
		return
	}
}
