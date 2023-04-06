// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registry

import (
	"examples/pkg/adapter/framework/gin/handler"
	"examples/pkg/adapter/infra"
	"examples/pkg/adapter/repository"
	"examples/pkg/logic"
)

// Injectors from app_container.go:

// wire用
func InitializeAppContainer(sqlh infra.SqlHandler, txh infra.TxHandler, store infra.LocalStore) *handler.AppContainer {
	userRepository := repository.NewUserRepository(sqlh)
	loginRepository := repository.NewLoginRepository(sqlh)
	transaction := repository.NewTransaction(txh)
	userLogic := logic.NewUserLogic(userRepository, loginRepository, transaction)
	userHandler := handler.NewUserHandler(userLogic)
	sessionRepository := repository.NewSessionRepository(store)
	sessionLogic := logic.NewSessionLogic(sessionRepository, loginRepository, transaction)
	sessionHandler := handler.NewSessionHandler(userLogic, sessionLogic)
	appContainer := handler.NewAppContainer(userHandler, sessionHandler)
	return appContainer
}
