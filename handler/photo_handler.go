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

type PhotoHandlerInterface interface {
	PhotoHandler(w http.ResponseWriter, r *http.Request)
}
type PhotoHandler struct {
	db    *sql.DB
	photo *model.Photo
}

func NewPhotohandler(db *sql.DB, photo *model.Photo) PhotoHandlerInterface {
	return &PhotoHandler{db: db, photo: photo}
}

func (p *PhotoHandler) PhotoHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	switch r.Method {
	case http.MethodGet:
		p.PhotoGetAll(w, r)
	case http.MethodPost:
		p.PhotoPost(w, r)
	case http.MethodPut:
		p.PhotoUpdate(w, r, id)
	case http.MethodDelete:
		p.PhotoDelete(w, r, id)
	}
}

func (p *PhotoHandler) PhotoGetAll(w http.ResponseWriter, r *http.Request) {
	var result = []model.Photo{}
	sqlGet := "Select * from photo;"
	rows, err := config.Db.Query(sqlGet)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&p.photo.Photo_id,
			&p.photo.Title,
			&p.photo.Caption,
			&p.photo.Photo_url,
			&p.photo.User_id,
			&p.photo.Created_at,
			&p.photo.Updated_at,
		); err != nil {
			fmt.Println("No Data", err)
		}
		result = append(result, *p.photo)
	}
	jsonData, _ := json.Marshal(&result)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
func (p *PhotoHandler) PhotoPost(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&p.photo)
	sqlSt := `insert into users (p_title, p_caption, p_url, p_created_date, p_updated_date)
		values ($1, $2, $3, $4, $5)
		returning p_id;`
	err := config.Db.QueryRow(sqlSt,
		p.photo.Title,
		p.photo.Caption,
		p.photo.Photo_url,
		time.Now(),
		time.Now(),
	).Scan(&p.photo.Photo_id)
	if err != nil {
		panic(err)
	}
	Register_respone := model.PhotoRegisterRespone{
		R_photo_id:   p.photo.Photo_id,
		R_title:      p.photo.Title,
		R_caption:    p.photo.Caption,
		R_photo_url:  p.photo.Photo_url,
		R_user_id:    p.photo.User_id,
		R_created_at: p.photo.Created_at,
	}
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

	fmt.Println(p.photo)
	fmt.Println(Register_respone)
	return
}
func (p *PhotoHandler) PhotoUpdate(w http.ResponseWriter, r *http.Request, id string) {
	for id != "" {

		//var orders = order.Order{}
		json.NewDecoder(r.Body).Decode(&p.photo)
		sqlSt := `update photo set
		p_title = $2, p_caption = $3, p_photo_url = $4, p_updated_date = $5 where p_id = $1;`
		res, err := config.Db.Exec(sqlSt,
			id,
			p.photo.Title,
			p.photo.Caption,
			p.photo.Photo_url,
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
		Update_respone := model.PhotoUpdateRespone{
			U_photo_id:   p.photo.Photo_id,
			U_title:      p.photo.Title,
			U_caption:    p.photo.Caption,
			U_photo_url:  p.photo.Photo_url,
			U_user_id:    p.photo.User_id,
			U_updated_at: p.photo.Updated_at,
		}
		jsonData, _ := json.Marshal(Update_respone)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		return
	}
}
func (p *PhotoHandler) PhotoDelete(w http.ResponseWriter, r *http.Request, id string) {
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
		w.Write([]byte(fmt.Sprint("Message: Your account has been successfully deleted ", count)))
		return
	}
}
