package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/logger"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	UserHandler          *handler.UserHandler
	TransactionHandler   *handler.TransactionHandler
	GameHandler          *handler.GameHandler
	ResetPasswordHandler *handler.ResetPasswordHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	router.Use(gin.Recovery())
	router.Use(middleware.Logger(logger.NewLogger()))
	router.Use(middleware.WithTimeout)
	router.Use(middleware.AuthorizeHandler())
	router.Use(middleware.ErrorHandler())

	users := router.Group("/users")
	users.GET("", opts.UserHandler.HandleGetUsers)
	users.POST("/register", opts.UserHandler.HandleUserRegister)
	users.POST("/login", opts.UserHandler.HandleUserLogin)
	users.POST("/reset-password", opts.ResetPasswordHandler.HandleRequestPassReset)
	users.PUT("/reset-password", opts.ResetPasswordHandler.HandleApplyPassReset)

	transactions := router.Group("/transactions")
	transactions.GET("", opts.TransactionHandler.HandleGetTransactions)
	transactions.POST("/top-up", opts.TransactionHandler.HandleTopUp)
	transactions.POST("/transfer", opts.TransactionHandler.HandleTransfer)

	router.GET("/user-details", opts.UserHandler.HandleGetUserDetails)

	games := router.Group("/games")
	games.GET("/boxes", opts.GameHandler.HandleGetAllBoxes)
	games.GET("/attempts", opts.GameHandler.HandleGetRemainingAttempt)
	games.POST("", opts.GameHandler.HandleChooseBox)

	return router
}
