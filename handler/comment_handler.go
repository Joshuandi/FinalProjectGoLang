package handler

import (
	"FinalProjectGoLang/config"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CommentHandlerInterface interface {
	CommentHandler(w http.ResponseWriter, r *http.Request)
}
type CommentHandler struct {
	db *sql.DB
}

func NewCommenthandler(db *sql.DB) CommentHandlerInterface {
	return &CommentHandler{db: db}
}

func (u *CommentHandler) CommentHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	switch r.Method {
	case http.MethodGet:
		u.CommentGetAll(w, r)
	case http.MethodPost:
		u.CommentPost(w, r)
	case http.MethodPut:
		u.CommentUpdate(w, r, id)
	case http.MethodDelete:
		u.CommentDelete(w, r, id)
	}
}

func (u *CommentHandler) CommentGetAll(w http.ResponseWriter, r *http.Request)            {}
func (u *CommentHandler) CommentPost(w http.ResponseWriter, r *http.Request)              {}
func (u *CommentHandler) CommentUpdate(w http.ResponseWriter, r *http.Request, id string) {}
func (u *CommentHandler) CommentDelete(w http.ResponseWriter, r *http.Request, id string) {
	sqlDelete := `DELETE from comment_ WHERE c_id = $1`
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
