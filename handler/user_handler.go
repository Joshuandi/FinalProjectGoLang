package handler

import (
	"FinalProjectGoLang/auth"
	"FinalProjectGoLang/config"
	"FinalProjectGoLang/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	switch r.Method {
	case http.MethodGet:
		u.UserGetAll(w, r)
	case http.MethodPost:
		u.UserPostLogin(w, r)
	case http.MethodPut:
		u.UserUpdate(w, r)
	case http.MethodDelete:
		u.UserDelete(w, r)
	}
}

func (u *UserHandler) UserGetAll(w http.ResponseWriter, r *http.Request) {
	var result = []model.User{}
	var users model.User
	sqlGet := "Select u_id, u_username, u_email, u_age from users;"
	rows, err := config.Db.Query(sqlGet)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&users.User_id,
			&users.Username,
			&users.Email,
			&users.Age,
		); err != nil {
			fmt.Println("No Data", err)
		}
		result = append(result, users)
	}
	jsonData, _ := json.Marshal(&result)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
func (u *UserHandler) UserPostRegister(w http.ResponseWriter, r *http.Request) {
	var users model.User
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
	fmt.Println("ini user", users)
	Register_respone := model.UserRegisterRespone{
		R_user_id:  users.User_id,
		R_email:    users.Email,
		R_username: users.Username,
		R_age:      users.Age,
	}
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

	fmt.Println("ini respone", Register_respone)
	return
	//}
}

func (u *UserHandler) UserPostLogin(w http.ResponseWriter, r *http.Request) {
	//var regist model.UserPostLogin
	var users model.User
	var login model.UserPostLogin
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	sqlSt := `select u_email, u_pass from users`
	rows, err := config.Db.Query(sqlSt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Println("ini log", login)
	for rows.Next() {
		if err = rows.Scan(
			&users.Email,
			&users.Password,
		); err != nil {
			fmt.Println("No Data", err)
		}
	}
	fmt.Println("ini us", users)
	check := auth.CheckHashPassword(login.Password, users.Password)
	if !check {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println("ini check", check)
	validToken, err := auth.GenerateJWT(users.Email, users.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println("ini valid token", validToken)
	var token model.UserToken
	token.TokenString = validToken
	fmt.Println("ini token", token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)

}
func (u *UserHandler) UserUpdate(w http.ResponseWriter, r *http.Request) {
	//var orders = order.Order{}
	var users model.User
	json.NewDecoder(r.Body).Decode(&users)
	sqlSt := `update users set
		u_username = $2, u_email = $3, u_updated_date = $4 where u_id = $1;`
	res, err := config.Db.Exec(sqlSt,
		id,
		users.Username,
		users.Email,
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
		U_user_id:    users.User_id,
		U_email:      users.Email,
		U_username:   users.Username,
		U_age:        users.Age,
		U_Updated_at: users.Updated_at,
	}
	jsonData, _ := json.Marshal(Update_respone)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	return
}

func (u *UserHandler) UserDelete(w http.ResponseWriter, r *http.Request) {
	sqlDelete := `DELETE from users WHERE u_id = $1`
	res, err := config.Db.Exec(sqlDelete)
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
