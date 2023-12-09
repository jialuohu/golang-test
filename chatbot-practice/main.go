package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)
import chatbot "golang-test-proj/chatbot-practice/base"

var chatBotName string

func init() {
	flag.StringVar(&chatBotName, "chatbot", "simple.en", "The chatbot's name for dialogue.")
}

func main() {
	flag.Parse()
	chatbot.Register(chatbot.NewEnglishBot(chatBotName, nil))
	myChatBot := chatbot.GetChatBot(chatBotName)
	if myChatBot == nil {
		err := fmt.Errorf("Fatal error: Unsupported chatbot named %s\n", chatBotName)
		checkError(nil, err, true)
	}

	begin, err := myChatBot.Begin()
	checkError(myChatBot, err, true)
	fmt.Println(begin)

	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	checkError(myChatBot, err, true)
	fmt.Println(myChatBot.Hello(input[:len(input)-1]))

	for {
		input, err := inputReader.ReadString('\n')
		if checkError(myChatBot, err, false) {
			continue
		}
		output, end, err := myChatBot.Talk(input)
		if checkError(myChatBot, err, false) {
			continue
		}
		if output != "" {
			fmt.Println(output)
		}
		if end {
			err = myChatBot.End()
			checkError(myChatBot, err, false)
			os.Exit(0)
		}
	}
}

func checkError(chatBot chatbot.ChatBot, err error, exit bool) bool {
	if err == nil {
		return false
	}
	if chatBot != nil {
		fmt.Println(chatBot.ReportError(err))
	} else {
		fmt.Println(err)
	}
	if exit {
		debug.PrintStack()
		os.Exit(1)
	}
	return true
}
