package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Despenrado/webMesk/internal/model"
	"github.com/Despenrado/webMesk/internal/service"
	"github.com/Despenrado/webMesk/pkg/utils"
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
		usr.Sanitize()
		utils.Respond(w, r, http.StatusCreated, usr)
	})
}

func (uh *UserHandler) FindUserById() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println("FindUserById")
		sid, ok := vars["id"]
		fmt.Println(vars)
		fmt.Println(sid)
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		fmt.Println("FindUserById")
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		usr, err := uh.service.User().FindById(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		usr.Sanitize()
		utils.Respond(w, r, http.StatusFound, usr)
	})
}

func (uh *UserHandler) ReadUsersLimitedList() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ReadUsersLimitedList")
		vars := r.URL.Query()
		skip, err := strconv.Atoi(vars.Get("skip"))
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		limit, err := strconv.Atoi(vars.Get("limit"))
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		users, err := uh.service.User().ReadAll(r.Context(), skip, limit)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		for i, _ := range users {
			users[i].Sanitize()
		}
		utils.Respond(w, r, http.StatusOK, users)
	})
}

func (uh *UserHandler) UpdateUserByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sid, ok := vars["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
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
		usr.Sanitize()
		utils.Respond(w, r, http.StatusOK, usr)
	})
}

func (uh *UserHandler) DeleteUserByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sid, ok := vars["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err = uh.service.User().Delete(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}
