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

type ChatHandler struct {
	service service.Service
}

func NewChatHandler(service service.Service) *ChatHandler {
	return &ChatHandler{
		service: service,
	}
}

func (ch *ChatHandler) CreateChat() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chat := &model.Chat{}
		err := json.NewDecoder(r.Body).Decode(chat)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err = chat.Validate()
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		fmt.Println(chat)
		chat, err = ch.service.Chat().Create(r.Context(), chat)
		if err != nil {
			if err == utils.ErrRecordAlreadyExists {
				utils.Error(w, r, http.StatusBadRequest, err)
				return
			}
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		chat.Sanitize()
		utils.Respond(w, r, http.StatusCreated, chat)
	})
}

func (ch *ChatHandler) FindChatById() http.HandlerFunc {
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
		chat, err := ch.service.Chat().FindById(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		chat.Sanitize()
		utils.Respond(w, r, http.StatusFound, chat)
	})
}

func (ch *ChatHandler) FilterChats() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filter := &model.ChatFilter{}
		err := json.NewDecoder(r.Body).Decode(filter)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		chats, err := ch.service.Chat().FilterChat(r.Context(), filter)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		for i, _ := range chats {
			chats[i].Sanitize()
		}
		utils.Respond(w, r, http.StatusOK, chats)
	})
}

func (ch *ChatHandler) UpdateChatByID() http.HandlerFunc {
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
		chat := &model.Chat{}
		err = json.NewDecoder(r.Body).Decode(chat)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err = chat.Validate()
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		chat.ID = uint(id)
		chat, err = ch.service.Chat().Update(r.Context(), chat)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		chat.Sanitize()
		utils.Respond(w, r, http.StatusOK, chat)
	})
}

func (ch *ChatHandler) DeleteChatByID() http.HandlerFunc {
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
		err = ch.service.Chat().Delete(r.Context(), uint(id))
		if err != nil {
			utils.Error(w, r, http.StatusInternalServerError, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}
