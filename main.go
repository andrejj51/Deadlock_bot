package main

import (
	"github.com/cortinico/telebot"
	"github.com/gocolly/colly"
)

// значение winrate по типу 46%
func winrate() string {
	var value string
	// Инициализация Colly
	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.7.0.0 Safari/537.36"

	// Настройка запроса
	c.OnHTML("span.shad2", func(e *colly.HTMLElement) {
		// Получение значения div

		value = e.Text

	})

	// Отправка запроса
	err := c.Visit("https://deadlocktracker.gg/player/71035446")
	if err != nil {
		panic(err)
	}
	return value
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
			answer = winrate()
		default:
			answer = "You typed " + mess
		}
		return answer, nil
	})
}
