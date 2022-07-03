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

type MessageHandler struct {
	service service.Service
}

func NewMessageHandler(service service.Service) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (mh *MessageHandler) CreateMessage() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := &model.Message{}
		err := json.NewDecoder(r.Body).Decode(message)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err = message.Validate()
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		fmt.Println(message)
		message, err = mh.service.Message().Create(r.Context(), message)
		if err != nil {
			if err == utils.ErrRecordAlreadyExists {
				utils.Error(w, r, http.StatusBadRequest, err)
				return
			}
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusCreated, message)
	})
}

func (mh *MessageHandler) FindMessageByID() http.HandlerFunc {
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
		message, err := mh.service.Message().FindById(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		message.Sanitize()
		utils.Respond(w, r, http.StatusFound, message)
	})
}

func (mh *MessageHandler) ReadMessagesLimitedList() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		message, err := mh.service.Message().ReadAll(r.Context(), skip, limit)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		for i, _ := range message {
			message[i].Sanitize()
		}
		utils.Respond(w, r, http.StatusOK, message)
	})
}

func (mh *MessageHandler) UpdateMessageByID() http.HandlerFunc {
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
		message := &model.Message{}
		err = json.NewDecoder(r.Body).Decode(message)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err = message.Validate()
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		message.ID = uint(id)
		message, err = mh.service.Message().Update(r.Context(), message)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		message.Sanitize()
		utils.Respond(w, r, http.StatusOK, message)
	})
}

func (mh *MessageHandler) DeleteMessageByID() http.HandlerFunc {
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
		err = mh.service.Message().Delete(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}
