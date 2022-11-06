package restapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/service"
	"github.com/Despenrado/webMesk/pkg/utils"
)

type AuthHandler struct {
	service service.Service
}

func NewAuthHandler(service service.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (ah *AuthHandler) SignUp() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		token, err := ah.service.Auth().SignUp(r.Context(), user)
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Authorization", "Bearer "+token)
		user.Sanitize()
		utils.Respond(w, r, http.StatusCreated, user)
	})
}

func (ah *AuthHandler) SignIn() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		user, token, err := ah.service.Auth().SignIn(r.Context(), user)
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Authorization", "Bearer "+token)
		user.Sanitize()
		utils.Respond(w, r, http.StatusCreated, user)
	})
}

func (ah *AuthHandler) Logout() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("missing authorization header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		err := ah.service.Auth().LogOut(r.Context(), tokenString)
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}

func (ah *AuthHandler) ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("missing authorization header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		ctx, err := ah.service.Auth().ValidAuthorization(r.Context(), tokenString)
		r = r.WithContext(ctx)
		if err != nil {
			utils.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		// log.Println(r.Context().Value("user_id"))
		next.ServeHTTP(w, r)
	})
}
