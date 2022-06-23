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

type SocialMediaHandlerInterface interface {
	SocialMediaHandler(w http.ResponseWriter, r *http.Request)
}
type SocialMediaHandler struct {
	db *sql.DB
	sm *model.SocialMedia
}

func NewSocialMediahandler(db *sql.DB, sm *model.SocialMedia) SocialMediaHandlerInterface {
	return &SocialMediaHandler{db: db, sm: sm}
}

func (s *SocialMediaHandler) SocialMediaHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	switch r.Method {
	case http.MethodGet:
		s.SocialMediaGetAll(w, r)
	case http.MethodPost:
		s.SocialMediaPost(w, r)
	case http.MethodPut:
		s.SocialMediaUpdate(w, r, id)
	case http.MethodDelete:
		s.SocialMediaDelete(w, r, id)
	}
}

func (s *SocialMediaHandler) SocialMediaGetAll(w http.ResponseWriter, r *http.Request) {
	var result = []model.SocialMedia{}
	sqlGet := "Select * from social_media;"
	rows, err := config.Db.Query(sqlGet)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&s.sm.Sm_Id,
			&s.sm.Name,
			&s.sm.SocialMedia_url,
			&s.sm.User_id,
			&s.sm.Created_at,
			&s.sm.Updated_at,
			&s.sm.User,
		); err != nil {
			fmt.Println("No Data", err)
		}
		result = append(result, *s.sm)
	}
	jsonData, _ := json.Marshal(&result)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
func (s *SocialMediaHandler) SocialMediaPost(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&s.sm)
	sqlSt := `insert into social_media (sm_name, sm_url, sm_created_at, user_id)
		values ($1, $2, $3, $4)
		returning sm_id;`
	err := config.Db.QueryRow(sqlSt,
		s.sm.Name,
		s.sm.SocialMedia_url,
		s.sm.Created_at,
		s.sm.User_id,
		s.sm.User,
	).Scan(&s.sm.Sm_Id)
	if err != nil {
		panic(err)
	}
	Register_respone := model.SocialMediaRegisterRespone{
		R_Sm_Id:           s.sm.Sm_Id,
		R_Name:            s.sm.Name,
		R_SocialMedia_url: s.sm.SocialMedia_url,
		R_User_id:         s.sm.User_id,
		R_Created_at:      s.sm.Created_at,
		R_User:            s.sm.User,
	}
	jsonData, _ := json.Marshal(Register_respone)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

	fmt.Println(s.sm)
	fmt.Println(Register_respone)
	return
}
func (s *SocialMediaHandler) SocialMediaUpdate(w http.ResponseWriter, r *http.Request, id string) {
	for id != "" {

		//var orders = order.Order{}
		json.NewDecoder(r.Body).Decode(&s.sm)
		sqlSt := `update social_media set
		sm_name = $2, sm_url = $3, u_updated_date = $4 where sm_id = $1;`
		res, err := config.Db.Exec(sqlSt,
			id,
			s.sm.Name,
			s.sm.SocialMedia_url,
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
		Update_respone := model.SocialMediaUpdateRespone{
			U_Sm_Id:           s.sm.Sm_Id,
			U_Name:            s.sm.Name,
			U_SocialMedia_url: s.sm.SocialMedia_url,
			U_User_id:         s.sm.User_id,
			U_Updated_at:      s.sm.Updated_at,
		}
		jsonData, _ := json.Marshal(Update_respone)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		return
	}
}
func (s *SocialMediaHandler) SocialMediaDelete(w http.ResponseWriter, r *http.Request, id string) {
	sqlDelete := `DELETE from social_media WHERE sm_id = $1`
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
