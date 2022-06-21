package service

import (
	"FinalProjectGoLang/model"
	"errors"
	"strings"
)

type ServiceInterface interface {
	UserValidation(users *model.User) (*model.User, error)
	PhotoValidation(photos *model.Photo) (*model.Photo, error)
	CommentValidation(comments *model.Comment) (*model.Comment, error)
	SocialMediaValidation(socialmedias *model.SocialMedia) (*model.SocialMedia, error)
}

type ServiceModel struct {
	users        *model.User
	photos       *model.Photo
	comments     *model.Comment
	socialmedias *model.SocialMedia
}

func NewService() ServiceInterface {
	return &ServiceModel{}
}

func (s *ServiceModel) UserValidation(users *model.User) (*model.User, error) {
	//emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if users.Email == "" {
		return nil, errors.New("Email harus di isi")
	}
	// if emailRegex.MatchString(user.Email) {
	// 	return nil, errors.New("Email harus sesuai")
	// }
	if !strings.Contains(users.Email, "@gmail.com") {
		return nil, errors.New("Must contain @gmail.com")
	}
	if users.Username == "" {
		return nil, errors.New("Password harus di isi")
	}
	if users.Password == "" || len(users.Password) < 6 {
		return nil, errors.New("Password harus di isi dan harus lebih dari 6 huruf")
	}
	if users.Age == 0 || users.Age <= 8 {
		return nil, errors.New("Umur harus di isi dan diatas 8 tahun")
	}
	return users, nil
}
func (s *ServiceModel) PhotoValidation(photos *model.Photo) (*model.Photo, error) {
	if photos.Title == "" {
		return nil, errors.New("Title harus di isi")
	}
	if photos.Photo_url == "" {
		return nil, errors.New("Photo Url harus di isi")
	}
	return photos, nil
}
func (s *ServiceModel) CommentValidation(comments *model.Comment) (*model.Comment, error) {
	if comments.Message == "" {
		return nil, errors.New("Comment harus di isi")
	}
	return comments, nil
}
func (s *ServiceModel) SocialMediaValidation(socialmedias *model.SocialMedia) (*model.SocialMedia, error) {
	if socialmedias.Name == "" {
		return nil, errors.New("Social Media name harus di isi")
	}
	if socialmedias.SocialMedia_url == "" {
		return nil, errors.New("Social Media Url harus di isi")
	}
	return socialmedias, nil
}
