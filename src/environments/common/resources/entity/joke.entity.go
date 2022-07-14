package entity

type Joke struct {
	ID    string `json:"id" groups:"client"`
	Value string `json:"value" groups:"client"`
	URL   string `json:"url" groups:"client"`
}
