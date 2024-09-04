package main

import (
	"fmt"

	"github.com/cortinico/telebot"
	"github.com/gocolly/colly"
)

type player struct {
	name      string
	url       string
	rate      string
	statistic []string
}

var playerUrl = map[string][2]string{
	"Андрей": [2]string{"Андрей", "https://deadlocktracker.gg/player/71035446"},
	"Макс":   [2]string{"Макс", "https://deadlocktracker.gg/player/202150072"},
}

/*func answer(players map[string]player) string{

}*/

// значение winrate по типу 46%
// принимает url player[name]
func winrate(playerUrlValue string) string {
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
	err := c.Visit(playerUrlValue)
	if err != nil {
		panic(err)
	}
	return value
}

func stats(playerUrlValue string) []string {
	var value string
	var list []string
	// Инициализация Colly
	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.7.0.0 Safari/537.36"

	// Настройка запроса
	c.OnHTML("span.\\32 shad2", func(e *colly.HTMLElement) {
		// Получение значения div

		value = e.Text
		list = append(list, value)

	})

	// Отправка запроса
	err := c.Visit(playerUrlValue)
	if err != nil {
		panic(err)
	}
	return list
}

// принимает ключ
// возвращает профиль
func profile(playerUrlValue string) player {

	return player{
		name: playerUrl[playerUrlValue][0],
		url:  playerUrl[playerUrlValue][1],
		rate: winrate(playerUrl[playerUrlValue][1]), // Winrate
		statistic: []string{
			stats(playerUrl[playerUrlValue][1])[0],  // Matches
			stats(playerUrl[playerUrlValue][1])[1],  // DLT Rating
			stats(playerUrl[playerUrlValue][1])[2],  // Max Rating
			stats(playerUrl[playerUrlValue][1])[3],  // KDA
			stats(playerUrl[playerUrlValue][1])[4],  // Souls/min
			stats(playerUrl[playerUrlValue][1])[5],  // Kills
			stats(playerUrl[playerUrlValue][1])[6],  // Creeps Kills
			stats(playerUrl[playerUrlValue][1])[7],  // Deaths
			stats(playerUrl[playerUrlValue][1])[8],  // Neutrals
			stats(playerUrl[playerUrlValue][1])[9],  // Assists
			stats(playerUrl[playerUrlValue][1])[10], // LastHits/min
			stats(playerUrl[playerUrlValue][1])[11], // AVG Damage
			stats(playerUrl[playerUrlValue][1])[12], // AVG Denies
		},
	}
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
			fmt.Println(profile("Андрей"))
		case "/deadlock":
			answer = "кек"
		default:
			answer = "You typed " + mess
		}
		return answer, nil
	})
}
