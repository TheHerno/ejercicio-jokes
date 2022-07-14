package interfaces

import (
	"net/http"
	"test-project-hernan/src/environments/common/resources/entity"
)

type IJokeService interface {
	GetJokes() ([]entity.Joke, error)
}

type IJokeController interface {
	GetJokes(response http.ResponseWriter, request *http.Request)
}
