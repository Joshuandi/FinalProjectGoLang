package handler

import (
	"FinalProjectGoLang/config"
	user "FinalProjectGoLang/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	sqlGet := "Select * from users;"
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
			&userss.Password,
			&userss.Age,
			&userss.Created_at,
			&userss.Updated_at,
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
