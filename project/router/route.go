package router

import (
	"github.com/e421083458/gateway_demo/project/controller"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	//设置cookie存储
	//https://github.com/gin-contrib/sessions
	store, _ := redis.NewStore(10, "tcp", lib.GetStringConf("redis_map.session.server"), "", []byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	})

	//admin_login
	adminLogin := router.Group("/admin_login")
	adminLogin.Use(
		sessions.Sessions("mysession", store),
		middleware.TranslationMiddleware(),
	)
	{
		controller.AdminLoginRegister(adminLogin)
	}

	//admin
	admin := router.Group("/admin")
	admin.Use(
		middleware.TranslationMiddleware(),
		sessions.Sessions("mysession", store),
		middleware.AdminSessionAuthMiddleware(),
	)
	{
		controller.AdminRegister(admin)
	}

	//service
	service := router.Group("/service")
	service.Use(
		middleware.TranslationMiddleware(),
		sessions.Sessions("mysession", store),
		middleware.AdminSessionAuthMiddleware(),
	)
	{
		controller.ServiceRegister(service)
	}

	//app
	app := router.Group("/app")
	app.Use(
		middleware.TranslationMiddleware(),
		sessions.Sessions("mysession", store),
		middleware.AdminSessionAuthMiddleware(),
	)
	{
		controller.APPRegister(app)
	}

	//app
	dash := router.Group("/dashboard")
	dash.Use(
		middleware.TranslationMiddleware(),
		sessions.Sessions("mysession", store),
		middleware.AdminSessionAuthMiddleware(),
	)
	{
		controller.DashBoardRegister(dash)
	}
	return router
}
