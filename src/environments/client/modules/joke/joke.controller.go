package joke

import (
	"net/http"
	"test-project-hernan/src/environments/client/resources/controller"
	"test-project-hernan/src/environments/client/resources/interfaces"
)

// struct that implements IJokeController
type jokeController struct {
	controller.ClientController
	jokeService interfaces.IJokeService
}

func NewJokeController(jService interfaces.IJokeService) interfaces.IJokeController {
	jokeController := jokeController{}
	jokeController.jokeService = jService
	return jokeController
}

/*
GetJokes returns a list of 25 jokes from the API
*/
func (j jokeController) GetJokes(response http.ResponseWriter, request *http.Request) {
	jokes, err := j.jokeService.GetJokes()
	if err != nil {
		j.MakeErrorResponse(response, err)
	}
	j.MakeSuccessResponse(response, jokes, http.StatusOK, "Jokes retrieved successfully")
}
