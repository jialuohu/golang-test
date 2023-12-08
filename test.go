package main

import "fmt"

type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type myTalk string

func (talk *myTalk) Hello(userName string) string {
	result := "Hello, " + userName + "\n"
	return result
}

func (talk *myTalk) Talk(heard string) (string, bool, error) {
	return "I heard you said: " + heard + "\n", true, nil
}

func (talk *myTalk) Name() string {
	return "Curly Anteater"
}

func (talk *myTalk) Begin() (string, error) {
	return "Started", nil
}

func (talk *myTalk) ReportError(err error) string {
	return "No error so far!"
}

func (talk *myTalk) End() error {
	return nil
}

func main() {
	var talk Talk = new(myTalk)
	str, ok := talk.(*myTalk)
	if ok == true {
		fmt.Println("OK, " + str.Hello("Dog"))
	} else {
		fmt.Println("Error")
	}

}
