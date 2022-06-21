package handler

import (
	"FinalProjectGoLang/config"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PhotoHandlerInterface interface {
	PhotoHandler(w http.ResponseWriter, r *http.Request)
}
type PhotoHandler struct {
	db *sql.DB
}

func NewPhotohandler(db *sql.DB) PhotoHandlerInterface {
	return &PhotoHandler{db: db}
}

func (u *PhotoHandler) PhotoHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	switch r.Method {
	case http.MethodGet:
		u.PhotoGetAll(w, r)
	case http.MethodPost:
		u.PhotoPost(w, r)
	case http.MethodPut:
		u.PhotoUpdate(w, r, id)
	case http.MethodDelete:
		u.PhotoDelete(w, r, id)
	}
}

func (u *PhotoHandler) PhotoGetAll(w http.ResponseWriter, r *http.Request)            {}
func (u *PhotoHandler) PhotoPost(w http.ResponseWriter, r *http.Request)              {}
func (u *PhotoHandler) PhotoUpdate(w http.ResponseWriter, r *http.Request, id string) {}
func (u *PhotoHandler) PhotoDelete(w http.ResponseWriter, r *http.Request, id string) {
	sqlDelete := `DELETE from photo WHERE p_id = $1`
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
