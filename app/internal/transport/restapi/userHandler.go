package restapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/service"
	"github.com/Despenrado/webMesk/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gopkg.in/gorilla/mux.v1"
)

type UserHandler struct {
	service service.Service
}

func NewUserHandler(service service.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) CreateUser() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usr := &model.User{}
		err := json.NewDecoder(r.Body).Decode(usr)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err = usr.Validate()
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		usr, err = uh.service.User().Create(r.Context(), usr)
		if err != nil {
			if err == utils.ErrRecordAlreadyExists {
				utils.Error(w, r, http.StatusBadRequest, err)
				return
			}
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, usr)
	})
}

func (uh *UserHandler) FindUserById() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sid, ok := vars["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		id, err := strconv.ParseUint(sid, 10, 32)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		usr, err := uh.service.User().FindById(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, usr)
	})
}

func (uh *UserHandler) FilterUsers() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filter := &model.UserFilter{}
		err := decoder.Decode(filter, r.URL.Query())
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		users, err := uh.service.User().FilterUser(r.Context(), filter)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, users)
	})
}

func (uh *UserHandler) UpdateUserByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.Context().Value("user_id").(string), 10, 32)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrUserNotFound)
			return
		}
		usr := &model.User{}
		err = json.NewDecoder(r.Body).Decode(usr)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err = usr.Validate()
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		usr.ID = uint(id)
		usr, err = uh.service.User().Update(r.Context(), usr)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, usr)
	})
}

func (uh *UserHandler) DeleteUserByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.Context().Value("user_id").(string), 10, 32)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrUserNotFound)
			return
		}
		err = uh.service.User().Delete(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}

func (uh *UserHandler) FindUserByEmail() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		email := vars.Get("email")
		if err := validation.Validate(email, is.Email); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		usr, err := uh.service.User().FindByEmail(r.Context(), email)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		usr.Sanitize()
		utils.Respond(w, r, http.StatusOK, usr)
	})
}
