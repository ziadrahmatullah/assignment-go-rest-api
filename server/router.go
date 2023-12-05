package server

// import "github.com/gin-gonic/gin"

// type RouterOpts struct {
// // 	JobHandler      *handler.JobHandler
// // 	UserHandler     *handler.UserHandler
// // 	ApplyJobHandler *handler.ApplyJobHandler
// }

// func NewRouter(opts RouterOpts) *gin.Engine {
// 	router := gin.New()
// 	router.ContextWithFallback = true

// 	// router.Use(gin.Recovery())
// 	// router.Use(middleware.WithTimeout)
// 	// router.Use(middleware.AuthorizeHandler())
// 	// router.Use(middleware.ErrorHandler())
// 	// router.Use(middleware.Logger(logger.NewLogger()))

// 	jobs := router.Group("/jobs")
// 	jobs.GET("", opts.JobHandler.HandleGetAllJobs)
// 	jobs.POST("", opts.JobHandler.HandleCreateJob)
// 	jobs.PUT("", opts.JobHandler.HandleUpdateJobExpireDate)
// 	jobs.DELETE("", opts.JobHandler.HandleDeleteJob)

// 	users := router.Group("/users")
// 	users.GET("", opts.UserHandler.HandleGetUsers)
// 	users.POST("/register", opts.UserHandler.HandleUserRegister)
// 	users.POST("/login", opts.UserHandler.HandleUserLogin)

// 	applyjobs := router.Group("/apply-jobs")
// 	applyjobs.GET("", opts.ApplyJobHandler.HandleGetAllRecords)
// 	applyjobs.POST("", opts.ApplyJobHandler.HandleCreateApplyJob)
// 	return router
// }
