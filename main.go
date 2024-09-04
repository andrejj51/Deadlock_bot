package main

import (
	"github.com/cortinico/telebot"
)

func Deadlock() string {
	return "228"
}

func main() {
	conf := telebot.Configuration{
		BotName: "Pechel_bot",
		ApiKey:  "7477752268:AAFlaObTt6OAyQSs_fhGUv05i13wGthtGxg"}

	var bot telebot.Bot

	bot.Start(conf, func(mess string) (string, error) {
		var answer string
		switch mess {
		case "/test":
			answer = "Test command works :)"
		case "/deadlock":
			answer = Deadlock()
		default:
			answer = "You typed " + mess
		}
		return answer, nil
	})
}
