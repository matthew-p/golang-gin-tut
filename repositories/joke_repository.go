package repositories

import (
	. "github.com/matthew-p/golang-gin-tut/models"
)

var jokes = []Joke{
	Joke{ID: 1, Likes: 0, Joke: "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{ID: 2, Likes: 0, Joke: "How many apples grow on a tree? All of them."},
	Joke{ID: 3, Likes: 0, Joke: "What do you call a fake noodle? An impasta."},
	Joke{ID: 4, Likes: 0, Joke: "Want to hear a joke about paper? Nevermind, it's tearable."},
	Joke{ID: 5, Likes: 0, Joke: "I just watched a program about beavers.It was the best dam program I've ever seen."},
	Joke{ID: 6, Likes: 0, Joke: "Why did the coffee file a police report? It got mugged."},
	Joke{ID: 7, Likes: 0, Joke: "How does a penguin build it's house? Igloos it together."},
}

// GetAllJokes Returns all available jokes as an array
func GetAllJokes() []Joke {
	return jokes
}
