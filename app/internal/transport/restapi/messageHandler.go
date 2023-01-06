package restapi

import (
	"encoding/json"
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
		userId, err := strconv.ParseUint(r.Context().Value("user_id").(string), 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrUserNotFound)
			return
		}
		message.UserID = uint(userId)
		err = message.Validate()
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		message, err = mh.service.Message().Create(r.Context(), message)
		if err != nil {
			if err == utils.ErrRecordAlreadyExists {
				utils.Error(w, r, http.StatusBadRequest, err)
				return
			}
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, message)
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
		userId, err := strconv.ParseUint(r.Context().Value("user_id").(string), 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrUserNotFound)
			return
		}
		message, err := mh.service.Message().FindById(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		if !message.CheckPermissions(uint(userId)) {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrNoPermissions)
		}
		utils.Respond(w, r, http.StatusOK, message)
	})
}

func (mh *MessageHandler) FilterMessages() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filter := &model.MessageFilter{}
		err := decoder.Decode(filter, r.URL.Query())
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		userId, err := strconv.ParseUint(r.Context().Value("user_id").(string), 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrUserNotFound)
			return
		}
		messages, err := mh.service.Message().FilterMessage(r.Context(), filter)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		messageList := []model.Message{}
		for _, v := range messages {
			if v.CheckPermissions(uint(userId)) {
				messageList = append(messageList, v)
			}
		}
		utils.Respond(w, r, http.StatusOK, messageList)
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

		userId, err := strconv.ParseUint(r.Context().Value("user_id").(string), 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrUserNotFound)
			return
		}
		message.ID = uint(id)
		message.UserID = uint(userId)
		message, err = mh.service.Message().Update(r.Context(), message)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
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

		userId, err := strconv.ParseUint(r.Context().Value("user_id").(string), 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrUserNotFound)
			return
		}
		message := &model.Message{
			ID:     uint(id),
			UserID: uint(userId),
		}

		err = mh.service.Message().Delete(r.Context(), message)
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}

func (mh *MessageHandler) MarkAsRead() http.HandlerFunc {
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
		usid, err := strconv.ParseUint(r.Context().Value("user_id").(string), 10, 64)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrUserNotFound)
			return
		}
		err = mh.service.Message().MarkAsRead(r.Context(), uint(id), uint(usid))
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}
