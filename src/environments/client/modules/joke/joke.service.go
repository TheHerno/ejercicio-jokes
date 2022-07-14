package joke

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"test-project-hernan/src/environments/client/resources/interfaces"
	"test-project-hernan/src/environments/common/resources/entity"
)

const URL = "https://api.chucknorris.io/jokes/random"

/*
Struct that implements IJokeService
*/
type jokeService struct {
}

func NewJokeService() interfaces.IJokeService {
	return jokeService{}
}

func getJoke(jokes map[string]entity.Joke, wg *sync.WaitGroup, lock *sync.RWMutex, errChan chan error) {
	getFromAPI := func() (*entity.Joke, error) {
		response, err := http.Get(URL)
		if err != nil {
			return nil, err
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		joke := &entity.Joke{}
		json.Unmarshal(responseData, joke)
		return joke, nil
	}

	done := false

	for !done {
		joke, err := getFromAPI()
		errChan <- err
		lock.Lock()
		if _, ok := jokes[joke.ID]; !ok {
			jokes[joke.ID] = *joke
			done = true
			wg.Done()
		}
		lock.Unlock()
	}
}

/*
GetJokes() returns a list of 25 jokes from the API
*/
func (j jokeService) GetJokes() ([]entity.Joke, error) {
	wg := sync.WaitGroup{}
	wg.Add(25)
	jokes := make(map[string]entity.Joke)
	errChan := make(chan error)
	var lock = sync.RWMutex{}
	for i := 0; i < 25; i++ {
		go getJoke(jokes, &wg, &lock, errChan)
	}
	for i := 0; i < 25; i++ {
		err := <-errChan
		if err != nil {
			return nil, err
		}
	}
	wg.Wait()
	jokesList := make([]entity.Joke, 0, len(jokes))
	for _, value := range jokes {
		jokesList = append(jokesList, value)
	}
	return jokesList, nil
}
