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

type UserRegisterRespone struct {
	R_user_id  int    `json:"user_id"`
	R_username string `json:"username"`
	R_email    string `json:"email"`
	R_age      int    `json:"age"`
}

type UserUpdateRespone struct {
	U_user_id    int       `json:"user_id"`
	U_email      string    `json:"email"`
	U_username   string    `json:"username"`
	U_age        int       `json:"age"`
	U_Updated_at time.Time `json:"updated_at"`
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

type PhotoRegisterRespone struct {
	R_photo_id   int       `json:"photo_id"`
	R_title      string    `json:"title"`
	R_caption    string    `json:"caption"`
	R_photo_url  string    `json:"photo_url"`
	R_user_id    int       `json:"User_id"`
	R_created_at time.Time `json:"create_at"`
}

type PhotoUpdateRespone struct {
	U_photo_id   int       `json:"photo_id"`
	U_title      string    `json:"title"`
	U_caption    string    `json:"caption"`
	U_photo_url  string    `json:"photo_url"`
	U_user_id    int       `json:"User_id"`
	U_updated_at time.Time `json:"updated_at"`
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

type CommentRegisterRespone struct {
	R_Comment_id int       `json:"comment_id"`
	R_User_id    int       `json:"User_id"`
	R_Photo_id   int       `json:"Photo_id"`
	R_Message    string    `json:"message"`
	R_Created_at time.Time `json:"create_at"`
}

type CommentUpdateRespone struct {
	U_Comment_id int       `json:"comment_id"`
	U_User_id    int       `json:"User_id"`
	U_Photo_id   int       `json:"Photo_id"`
	U_Message    string    `json:"message"`
	U_Updated_at time.Time `json:"updated_at"`
}

type SocialMedia struct {
	Sm_Id           int       `json:"socialMedia_id"`
	Name            string    `json:"name"`
	SocialMedia_url string    `json:"socialMedia_url"`
	Created_at      time.Time `json:"create_at"`
	Updated_at      time.Time `json:"updated_at"`
	User_id         int       `json:"User_id"`
	User            User      `json:"User"`
}

type SocialMediaRegisterRespone struct {
	R_Sm_Id           int       `json:"socialMedia_id"`
	R_Name            string    `json:"name"`
	R_SocialMedia_url string    `json:"socialMedia_url"`
	R_Created_at      time.Time `json:"create_at"`
	R_User_id         int       `json:"User_id"`
	R_User            User      `json:"User"`
}

type SocialMediaUpdateRespone struct {
	U_Sm_Id           int       `json:"socialMedia_id"`
	U_Name            string    `json:"name"`
	U_SocialMedia_url string    `json:"socialMedia_url"`
	U_Updated_at      time.Time `json:"updated_at"`
	U_User_id         int       `json:"User_id"`
	U_User            User      `json:"User"`
}
