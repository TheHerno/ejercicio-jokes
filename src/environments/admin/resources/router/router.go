package router

import (
	"test-project-hernan/src/libs/middleware"

	"github.com/gorilla/mux"
)

/*
SetupAdminRoutes creates all instances for admin enviroment and calls each router
*/
func SetupAdminRoutes(subRouter *mux.Router) {
	subRouter.Use(middleware.NewAuthMiddleware().HandlerAdmin())
}
