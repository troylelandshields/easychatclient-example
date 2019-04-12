package main

import (
	"fmt"

	"github.com/troylelandshields/easychat"
	"github.com/troylelandshields/easyinput"
)

func main() {

	fmt.Print("Enter your name: ")
	name := easyinput.TakeInput()
	if name == "" {
		return
	}

	chatClient := easychat.JoinChatRoom("localhost", name)

	go receiveMessagesLoop(chatClient)

	sendMessagesLoop(chatClient)
}

func receiveMessagesLoop(chatClient *easychat.ChatClient) {
	for {
		msg, ok := chatClient.ReceiveMessage()
		if !ok {
			return
		}

		fmt.Printf("\n[%s] %s: %s\n...> ",
			msg.Time.Format("Jan 2, 3:04:05 PM"), msg.From, msg.Body)
	}
}

func sendMessagesLoop(chatClient *easychat.ChatClient) {
	for {
		fmt.Print("> ")
		msg := easyinput.TakeInput()

		if msg == "" {
			continue
		}

		chatClient.SendMessage(msg)
	}
}
