package joke

import (
	"net/http"
	"test-project-hernan/src/environments/client/resources/interfaces"

	"github.com/gorilla/mux"
)

type jokeRouter struct {
	JokeController interfaces.IJokeController
}

func NewJokeRouter(subRouter *mux.Router, cJoke interfaces.IJokeController) {
	routerProduct := jokeRouter{cJoke}
	routerProduct.routes(subRouter)
}

/*
routes assigns controller function for routes
*/
func (r *jokeRouter) routes(subRouter *mux.Router) {
	subRouter.
		Path("").
		Handler(http.HandlerFunc(r.JokeController.GetJokes)).
		Methods(http.MethodGet)
}
