package restapi

import (
	"context"
	"net/http"

	"github.com/Despenrado/webMesk/internal/utils"
	pkgutils "github.com/Despenrado/webMesk/pkg/utils"
	"github.com/gorilla/schema"
	"go.uber.org/fx"
	"gopkg.in/gorilla/mux.v1"
)

// type Server struct {
// 	ctx            context.Context
// 	port           string
// 	logger         *utils.Logger
// 	router         *mux.Router
// 	userHandler    *UserHandler
// 	messageHandler *MessageHandler
// 	chatHandler    *ChatHandler
// 	authHandler    *AuthHandler
// }
var decoder = schema.NewDecoder()

func NewServer(
	lc fx.Lifecycle,
	logger *pkgutils.Logger,
	config *utils.Config,
) *mux.Router {
	// srv := &Server{
	// 	ctx:            ctx,
	// 	port:           port,
	// 	logger:         logger,
	// 	userHandler:    userHandler,
	// 	messageHandler: messageHandler,
	// 	chatHandler:    chatHandler,
	// 	authHandler:    authHandler,
	// 	router: ,
	// }
	router := mux.NewRouter()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Infof("Server started on port %s...", config.RestAPIServer.Port)
			go func() {
				if err := http.ListenAndServe(":"+config.RestAPIServer.Port, router); err != nil {
					logger.Warn("Server stopped", err)
				}
			}()
			return nil
		},
		// OnStop: func(ctx context.Context) error {
		// 	logger.Infof("Server stopped")
		// 	return server.Shutdown(ctx)
		// },
	})

	return router
}

// func (srv *Server) Run() error {
// 	defer srv.logger.Infof("Server stopped")
// 		srv.logger.Infof("Server started on port %s...\n", srv.port)
// 	if err := http.ListenAndServe(":"+srv.port, srv.router); err != nil {
// 		return err
// 	}
// 	return nil
// }

func RegisterHundlers(
	router *mux.Router,
	logger *pkgutils.Logger,
	userHandler *UserHandler,
	messageHandler *MessageHandler,
	chatHandler *ChatHandler,
	authHandler *AuthHandler,
) {
	pfx := "restapi"
	pfx = "/" + pfx
	router.Schemes("http")
	router.Use(pkgutils.SetRequestId)
	router.Use(logger.LogRequest)

	router.HandleFunc(pfx+"/signup", authHandler.SignUp()).Methods("POST")
	router.HandleFunc(pfx+"/signin", authHandler.SignIn()).Methods("PUT")

	authorizedRouter := router.NewRoute().Subrouter()
	authorizedRouter.Use(authHandler.ValidateToken)

	authorizedRouter.HandleFunc(pfx+"/logout", authHandler.Logout()).Methods("GET")

	usrRouter := authorizedRouter.PathPrefix(pfx + "/users").Subrouter()
	usrRouter.HandleFunc("", userHandler.CreateUser()).Methods("POST")
	usrRouter.HandleFunc("/filter", userHandler.FilterUsers()).Methods("GET")
	usrRouter.HandleFunc("/{id:[0-9]+}", userHandler.FindUserById()).Methods("GET")
	usrRouter.HandleFunc("", userHandler.UpdateUserByID()).Methods("PUT")
	usrRouter.HandleFunc("", userHandler.DeleteUserByID()).Methods("DELETE")

	chatRouter := authorizedRouter.PathPrefix(pfx + "/chats").Subrouter()
	chatRouter.HandleFunc("", chatHandler.CreateChat()).Methods("POST")
	chatRouter.HandleFunc("/filter", chatHandler.FilterChats()).Methods("GET")
	chatRouter.HandleFunc("/{id:[0-9]+}", chatHandler.FindChatById()).Methods("GET")
	chatRouter.HandleFunc("/{id:[0-9]+}", chatHandler.UpdateChatByID()).Methods("PUT")
	chatRouter.HandleFunc("/{id:[0-9]+}", chatHandler.DeleteChatByID()).Methods("DELETE")

	messageRouter := authorizedRouter.PathPrefix(pfx + "/messages").Subrouter()
	messageRouter.HandleFunc("", messageHandler.CreateMessage()).Methods("POST")
	messageRouter.HandleFunc("/filter", messageHandler.FilterMessages()).Methods("GET")
	messageRouter.HandleFunc("/{id:[0-9]+}", messageHandler.FindMessageByID()).Methods("GET")
	messageRouter.HandleFunc("/{id:[0-9]+}", messageHandler.UpdateMessageByID()).Methods("PUT")
	messageRouter.HandleFunc("/{id:[0-9]+}", messageHandler.DeleteMessageByID()).Methods("DELETE")
	messageRouter.HandleFunc("/markasread/{id:[0-9]+}", messageHandler.MarkAsRead()).Methods("PUT")
}
