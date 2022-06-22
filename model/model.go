package model

import (
	"time"
)

type User struct {
	User_id    int       `json:"user_id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Age        int       `json:"age"`
	Created_at time.Time `json:"create_at"`
	Updated_at time.Time `json:"updated_at"`
}

type RegisterRespone struct {
	R_user_id  int    `json:"user_id"`
	R_username string `json:"username"`
	R_email    string `json:"email"`
	R_age      int    `json:"age"`
}

type Photo struct {
	Photo_id   int       `json:"photo_id"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	Photo_url  string    `json:"photo_url"`
	User_id    int       `json:"User_id"`
	Created_at time.Time `json:"create_at"`
	Updated_at time.Time `json:"updated_at"`
	User       User      `json:"User"`
}

type Comment struct {
	Comment_id int       `json:"comment_id"`
	User_id    int       `json:"User_id"`
	Photo_id   int       `json:"Photo_id"`
	Message    string    `json:"message"`
	Created_at time.Time `json:"create_at"`
	Updated_at time.Time `json:"updated_at"`
	User       User      `json:"User"`
	Photo      Photo     `json:"Photo"`
}
type SocialMedia struct {
	Sm_Id           int    `json:"socialMedia_id"`
	Name            string `json:"name"`
	SocialMedia_url string `json:"socialMedia_url"`
	User_id         int    `json:"User_id"`
	User            User   `json:"User"`
}
