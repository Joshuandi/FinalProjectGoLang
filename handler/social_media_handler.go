package handler

import (
	"FinalProjectGoLang/config"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SocialMediaHandlerInterface interface {
	SocialMediaHandler(w http.ResponseWriter, r *http.Request)
}
type SocialMediaHandler struct {
	db *sql.DB
}

func NewSocialMediahandler(db *sql.DB) SocialMediaHandlerInterface {
	return &SocialMediaHandler{db: db}
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

func (s *SocialMediaHandler) SocialMediaGetAll(w http.ResponseWriter, r *http.Request)            {}
func (s *SocialMediaHandler) SocialMediaPost(w http.ResponseWriter, r *http.Request)              {}
func (s *SocialMediaHandler) SocialMediaUpdate(w http.ResponseWriter, r *http.Request, id string) {}
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
