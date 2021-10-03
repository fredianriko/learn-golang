package main

import (
	"fmt"

	"log"

	"example.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")// setting prefix for output
	log.SetFlags(0)// Setflag(0) is hiding the date and time of log

	//request greetings message
	// pass 2 input
	// log is buildin package to logging event or something

	// make new array containing "names"
	names := []string{"Gladys", "Samantha", "Darrin"}
   
	//invoke the greetings package and called the Hellos function from that package and put names as argument to that function
	message, err :=  greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

}


// names  := []string{"name1", "name2", "name3"}

//func Hello(name string) (string, err){
// 		if name == ""{
//     return "", err.New("empty string")
// 		message := fmt.Sprintf(randomGreetings(),name)
// 		return message, nil	
//}
//}

// func Hellos()



// func randomGreetings() string {
//  formats := []string{"Grettings 1 %v", "Greetings2 %v", "Greetings3 %v"}
//  return formats(rand.intn(len(formats)))
//}