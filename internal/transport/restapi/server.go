package restapi

import (
	"context"
	"net/http"

	"github.com/Despenrado/webMesk/pkg/utils"
	"gopkg.in/gorilla/mux.v1"
)

type Server struct {
	ctx            context.Context
	port           string
	logger         *utils.Logger
	router         *mux.Router
	userHandler    *UserHandler
	messageHandler *MessageHandler
	chatHandler    *ChatHandler
	authHandler    *AuthHandler
}

func NewServer(
	ctx context.Context,
	port string,
	ro *mux.Router,
	logger *utils.Logger,
	userHandler *UserHandler,
	messageHandler *MessageHandler,
	chatHandler *ChatHandler,
	authHandler *AuthHandler,
) *Server {
	srv := &Server{
		ctx:            ctx,
		port:           port,
		logger:         logger,
		userHandler:    userHandler,
		messageHandler: messageHandler,
		chatHandler:    chatHandler,
		authHandler:    authHandler,
	}
	if ro == nil {
		srv.InitDefaultEndpoints("/restapi")
	}
	return srv
}

func (srv *Server) Run() error {
	defer srv.logger.Infof("Server stopped")
	srv.logger.Infof("Server started on port %s...\n", srv.port)
	if err := http.ListenAndServe(":"+srv.port, srv.router); err != nil {
		return err
	}
	return nil
}

func (srv *Server) InitDefaultEndpoints(pfx string) {
	srv.router = mux.NewRouter()
	srv.router.Schemes("http")
	pfx = "/" + pfx
	// srv.router.Use(utils.SetRequestId)
	// srv.router.Use(srv.logger.LogRequest)

	usrRouter := srv.router.PathPrefix(pfx + "/users").Subrouter()
	usrRouter.HandleFunc("", srv.userHandler.CreateUser()).Methods("POST")
	usrRouter.HandleFunc("/read", srv.userHandler.ReadUsersLimitedList()).Methods("GET")
	usrRouter.HandleFunc("/{id:[0-9]+}", srv.userHandler.FindUserById()).Methods("GET")
	usrRouter.HandleFunc("/{id:[0-9]+}", srv.userHandler.UpdateUserByID()).Methods("PUT")
	usrRouter.HandleFunc("/{id:[0-9]+}", srv.userHandler.DeleteUserByID()).Methods("DELETE")
}
