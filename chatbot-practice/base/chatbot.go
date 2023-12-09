package base

import "errors"

type ChatBot interface {
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

var (
	ErrInvalidChatBotName = errors.New("invalid ChatBot name")
	ErrInvalidChatBot     = errors.New("invalid ChatBot")
	ErrExistingChatBot    = errors.New("existing ChatBot")
)

var chatBotMap = make(map[string]ChatBot, 10)

func Register(chatBot ChatBot) error {
	// Test if the chatBot passed in is nil
	if chatBot == nil {
		return ErrInvalidChatBot
	}

	// Test if the chatBot's name is empty (also can check its name is valid)
	name := chatBot.Name()
	if name == "" {
		return ErrInvalidChatBotName
	}

	// Test if the chatBot was already inserted into map
	if _, ok := chatBotMap[name]; ok {
		return ErrExistingChatBot
	}

	// Insert new chatBot
	chatBotMap[name] = chatBot
	return nil
}

func GetChatBot(name string) ChatBot {
	return chatBotMap[name]
}
