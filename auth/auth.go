package auth

import (
	"FinalProjectGoLang/config"
	"log"
	"net/http"
	"strings"

	"google.golang.org/api/idtoken"
)

type AuthMiddlewareInterface interface {
	AuthLoginValidation(next http.Handler) http.Handler
	AuthTokenValidation(next http.Handler) http.Handler
}
type AuthMiddleware struct {
	cfg config.Config
}

func NewAuthMiddleware(cfg *config.Config) AuthMiddlewareInterface {
	return &AuthMiddleware{cfg: *cfg}
}

func (a *AuthMiddleware) AuthLoginValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if uname, pass, ok := r.BasicAuth(); !ok || uname != a.cfg.Lusername || pass != a.cfg.Lpassword {
			w.WriteHeader(401)
			w.Write([]byte("ERROR DATA"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *AuthMiddleware) AuthTokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authhandler := strings.Split(r.Header.Get("Authorization"), " "); len(authhandler) == 2 && authhandler[0] == "Bearer" {
			idToken := authhandler[1]
			idTokenPayload, err := idtoken.Validate(r.Context(), idToken, a.cfg.GoogleClientId)
			if err != nil {
				log.Println("Error when validate id token", err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorixed Request"))
				return
			}
			log.Println(idTokenPayload)
			next.ServeHTTP(w, r)
		} else {
			log.Println("Error when validate id token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorixed Request"))
			return
		}
	})
}
