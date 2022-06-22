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
)

type CommentHandlerInterface interface {
	CommentHandler(w http.ResponseWriter, r *http.Request)
}
type CommentHandler struct {
	db      *sql.DB
	comment *model.Comment
}

func NewCommenthandler(db *sql.DB, comment *model.Comment) CommentHandlerInterface {
	return &CommentHandler{db: db, comment: comment}
}

func (c *CommentHandler) CommentHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	switch r.Method {
	case http.MethodGet:
		c.CommentGetAll(w, r)
	case http.MethodPost:
		c.CommentPost(w, r)
	case http.MethodPut:
		c.CommentUpdate(w, r, id)
	case http.MethodDelete:
		c.CommentDelete(w, r, id)
	}
}

func (c *CommentHandler) CommentGetAll(w http.ResponseWriter, r *http.Request) {
	var result = []model.Comment{}
	sqlGet := "Select * from comment;"
	rows, err := config.Db.Query(sqlGet)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&c.comment.Comment_id,
			&c.comment.Message,
			&c.comment.Photo_id,
			&c.comment.User_id,
			&c.comment.Updated_at,
			&c.comment.Created_at,
			&c.comment.User,
			&c.comment.Photo,
		); err != nil {
			fmt.Println("No Data", err)
		}
		result = append(result, *c.comment)
	}
	jsonData, _ := json.Marshal(&result)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
func (c *CommentHandler) CommentPost(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&c.comment.Comment_id)
	sqlSt := `insert into comment (c_message, photo_id, c_created_at)
		values ($1, $2, $3)
		returning c_id;`
	err := config.Db.QueryRow(sqlSt,
		c.comment.Message,
		c.comment.Photo_id,
		time.Now(),
	).Scan(&c.comment.Comment_id)
	if err != nil {
		panic(err)
	}
	Register_respone := model.CommentRegisterRespone{
		R_Comment_id: c.comment.Comment_id,
		R_Message:    c.comment.Message,
		R_Photo_id:   c.comment.Photo_id,
		R_User_id:    c.comment.User_id,
		R_Created_at: c.comment.Created_at,
	}
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

	fmt.Println(c.comment)
	fmt.Println(Register_respone)
	return
}
func (c *CommentHandler) CommentUpdate(w http.ResponseWriter, r *http.Request, id string) {
	for id != "" {

		//var orders = order.Order{}
		json.NewDecoder(r.Body).Decode(&c.comment)
		sqlSt := `update comment set c_message = $2, c_updated_date = $3 where c_id = $1;`
		res, err := config.Db.Exec(sqlSt,
			id,
			c.comment.Message,
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
		Update_respone := model.CommentUpdateRespone{
			U_Comment_id: c.comment.Comment_id,
			U_Message:    c.comment.Message,
			U_Photo_id:   c.comment.Photo_id,
			U_User_id:    c.comment.Photo_id,
			U_Updated_at: c.comment.Updated_at,
		}
		jsonData, _ := json.Marshal(Update_respone)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		return
	}
}
func (c *CommentHandler) CommentDelete(w http.ResponseWriter, r *http.Request, id string) {
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
		w.Write([]byte(fmt.Sprint("Message: Your account has been successfully deleted ", count)))
		return
	}
}
