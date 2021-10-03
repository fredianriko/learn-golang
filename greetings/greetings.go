package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello function
func Hello(name string) (string,error) {
	if name == "" {
		return "",errors.New("empty name")
	}

	// message := fmt.Sprintf(randomFormat(), name)

	message := fmt.Sprintf(randomFormat())
	return message,nil
}


// Hello function that accept map of string and return also map of string, 
func Hellos(names []string) (map[string]string, error){


	//make new "map" with name of messages
	messages := make(map[string]string)
	// loop over names map, and assign every item from the names map, into a variabel called "name"
	// and assign "message" vairiabel with Hello(name), this much more like a .map() function in javascriptr
	for _, name := range names{
		message, err := Hello(name)
		// if error not null, or the error is exist, then return nil as value of error
		if err != nil {
			return nil, err
		}

		//then assign messages
		messages[name] = message
	}
	return messages, nil
}

// init() function will only be called once per package call, so whenever you called a package with init() function, the task inside the init() function will always run first
func init () {
	rand.Seed(time.Now().UnixNano())
	// fmt.Println("test init function")
}


//random text
// this function returning each "greetings" based on the randomize index number from rand.Intn(len(format)), it randomize number in range of the format map length
func randomFormat() string{
	formats := []string{
		"Hi %v. Welcome",
		"Great to see you %v",
		"Hail, %v! Well met",
	}
	return formats[rand.Intn(len(formats))]
}
