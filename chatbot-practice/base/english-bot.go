package base

import (
	"fmt"
	"strings"
)

type englishBot struct {
	name string
	talk Talk
}

func (chatBot *englishBot) Name() string {
	return chatBot.name
}

func (chatBot *englishBot) Begin() (string, error) {
	return "Please input your name: ", nil
}

func (chatBot *englishBot) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	// With using "talk", we can now check if the chatBot has customized content
	if chatBot.talk != nil {
		return chatBot.talk.Hello(userName)
	}
	return fmt.Sprintf("Hello, %s! What can I do for you?", userName)
}

func (chatBot *englishBot) Talk(heard string) (string, bool, error) {
	heard = strings.TrimSpace(heard)
	if chatBot.talk != nil {
		return chatBot.talk.Talk(heard)
	}
	switch heard {
	case "":
		return "", false, nil
	case "nothing", "gdgd":
		return "Curly Seongjin", true, nil
	default:
		return "Sorry, I do not get it.", false, nil
	}
}

func (chatBot *englishBot) ReportError(err error) string {
	return fmt.Sprintf("An error occurred: %s\n", err)
}

func (chatBot *englishBot) End() error {
	return nil
}

func NewEnglishBot(name string, talk Talk) ChatBot {
	return &englishBot{
		name: name,
		talk: talk,
	}
}
