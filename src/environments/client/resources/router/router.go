package router

import (
	"test-project-hernan/src/environments/client/modules/joke"

	"github.com/gorilla/mux"
)

/*
SetupClientRoutes creates all instances for client enviroment and calls each router
*/
func SetupClientRoutes(subRouter *mux.Router) {
	jokeRoutes(subRouter.PathPrefix("/jokes").Subrouter())
}

/*
jokeRoutes sets the routes for the joke module
*/
func jokeRoutes(subRouter *mux.Router) {
	jService := joke.NewJokeService()
	jController := joke.NewJokeController(jService)
	joke.NewJokeRouter(subRouter, jController)
}
