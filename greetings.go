package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used
// in the rand package

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages
// Rhe returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hello there %v, nice to see you",
		"Hail, %v Well met!",
	}
	// return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
