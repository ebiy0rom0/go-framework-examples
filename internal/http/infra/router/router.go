package router

import (
	"examples/internal/http/infra/cache"
	"examples/internal/http/infra/middleware"
	"examples/internal/http/infra/web"
	"examples/internal/http/interface/handler"
	"examples/internal/http/interface/infra"
	"examples/internal/http/interface/repository"
	"examples/internal/http/logic"
	"net/http"
	"os"
)

func SetRoute(sqlh infra.SqlHandler) {
	// repository
	loginRepo := repository.NewLoginRepository(sqlh)
	userRepo := repository.NewUserRepository(sqlh)
	sessionRepo := repository.NewSessionRepository(cache.NewLocalStore())

	// logic
	userLogic := logic.NewUserLogic(userRepo)
	sessionLogic := logic.NewSessionLogic(sessionRepo, loginRepo)

	// middleware
	loggerMid := middleware.NewLoggerMiddleware(os.Stdout)
	authMid := middleware.NewAuthMiddleware(sessionRepo)

	// handler
	users := handler.NewUserHandler(userLogic)
	{
		http.Handle("/signup", loggerMid.WithLogger(web.HttpHandler(users.Signup)))
		http.Handle("/users/", loggerMid.WithLogger(web.HttpHandler(authMid.CheckToken(users.HandleRoot))))
	}

	session := handler.NewSessionHandler(userLogic, sessionLogic)
	{
		// サインアップ、サインインはトークン取得前なのでチェックを行わない
		http.Handle("/signin", loggerMid.WithLogger(web.HttpHandler(session.Signin)))
		http.Handle("/signout", loggerMid.WithLogger(web.HttpHandler(authMid.CheckToken(session.Signout))))
	}
}
