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
	"Андрей": {"Андрей", "https://deadlocktracker.gg/player/71035446"},
	"Макс":   {"Макс", "https://deadlocktracker.gg/player/202150072"},
}

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

	// Добавляем обработчик для ошибки
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("Error: %v\n", err)
	})

	// Отправка запроса
	err := c.Visit(playerUrlValue)
	if err != nil {
		panic(err)
	}
	if len(list) != 13 {
		fmt.Println("НЕ РАВЕН 13*************")
		fmt.Println(len(list))

	}
	return list

}

// принимает ключ
// возвращает профиль
func profile(playerUrlValue string) player {
	var player player
	player.name = playerUrl[playerUrlValue][0]
	player.url = playerUrl[playerUrlValue][1]
	player.rate = winrate(playerUrl[playerUrlValue][1])[2:] // Winrate
	player.statistic = stats(playerUrl[playerUrlValue][1])

	return player
}

// возвращает готовый ответ по профилю
func answerStat(player player) string {
	mes := "%s\n" +
		"Winrate: %s      | Matches: %s" +
		"DLT Raiting: %s  | Max Raiting: %s\n" +
		"KDA: %s          | Souls/min: %s\n" +
		"Kills: %s        | Creeps Kill: %s\n" +
		"Deaths: %s       | Naturals: %s\n" +
		"Assists: %s      | LastHits/min: %s\n" +
		"AVG Damage: %s   | AVG Denies: %s\n"

	return fmt.Sprintf(mes, player.name,
		player.rate, player.statistic[0],
		player.statistic[1], player.statistic[2],
		player.statistic[3], player.statistic[4],
		player.statistic[5], player.statistic[6],
		player.statistic[7], player.statistic[8],
		player.statistic[9], player.statistic[10],
		player.statistic[11], player.statistic[12],
	)
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
			// answer = answerStat(profile(playerUrl["Андрей"][0]))
			answer = answerStat(profile(playerUrl["Андрей"][0]))
		case "/deadlock":
			answer = "кек"
		default:
			answer = "You typed " + mess
		}
		return answer, nil
	})
}
