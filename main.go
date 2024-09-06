package main

import (
	"fmt"
	"sort"

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
	"Андрей":  {"Андрей", "https://deadlocktracker.gg/player/71035446"},
	"Макс":    {"Макс", "https://deadlocktracker.gg/player/202150072"},
	"Тимофей": {"Тимофей", "https://deadlocktracker.gg/player/1217462833"},
	"Димасик": {"Димасик", "https://deadlocktracker.gg/player/1713203171"},
	"Лёха":    {"Лёха", "https://deadlocktracker.gg/player/93910342"},
	"Саша":    {"Саша", "https://deadlocktracker.gg/player/97066121"},
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
	return list

}

// принимает ключ
// возвращает профиль
func profile(playerUrlValue string) player {
	var player player
	if len(stats(playerUrl[playerUrlValue][1])) == 0 {
		player.name = playerUrl[playerUrlValue][0]
		player.url = playerUrl[playerUrlValue][1]
		return player
	}
	player.name = playerUrl[playerUrlValue][0]
	player.url = playerUrl[playerUrlValue][1]
	player.rate = winrate(playerUrl[playerUrlValue][1])[2:] // Winrate
	player.statistic = stats(playerUrl[playerUrlValue][1])

	return player
}

// возвращает готовый ответ по профилю
func answerStat(player player) string {
	var mes string
	if player.rate == "" {
		mes = fmt.Sprintf("%s не открыл свою стату (попуск)", player.name)
		return mes
	}

	mes = "%s\n" +
		"Winrate: %s\nMatches: %s\n" +
		"DLT Raiting: %s\nMax Raiting: %s\n" +
		"KDA: %s\nSouls/min: %s\n" +
		"Kills: %s\nCreeps Kill: %s\n" +
		"Deaths: %s\nNaturals: %s\n" +
		"Assists: %s\nLastHits/min: %s\n" +
		"AVG Damage: %s\nAVG Denies: %s\n"

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

/*
func space(mes string) string {
	//var lenListString []int
	var line1 []string
	var line2 []string
	var line []string

	list := strings.Split(mes, "\n")

	for i := range list[:len(list)-2] {
		fmt.Println(list[i+1])
		line1 = append(line1, strings.Split(list[i+1], "|")[0])
		line2 = append(line2, strings.Split(list[i+1], "|")[1])
	}

	for i := 0; i < 7; i++ {
		line = append(line, fmt.Sprintf("%-28s"+"| "+"%s\n", line1[i], line2[i]))
	}
	result := strings.Join(line, "")
	return fmt.Sprintln(list[0] + "\n" + result)

}*/

func topWinrate() {
	var top map[string]player
	//var topSort map[string]string
	top = make(map[string]player)
	for i := range playerUrl {
		if profile(playerUrl[i][0]).name != "" {
			top[profile(playerUrl[i][0]).name] = profile(playerUrl[i][0])
		}
	}

	// сортировка
	/*topSort = make(map[string]string)
	for key, val := range top {
		topSort[key] = val.rate
	}*/

	values := make([]string, 0, len(top))
	for _, v := range top {
		values = append(values, v.rate)
	}
	sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })
	//fmt.Println(values)
	for _, k := range values {
		fmt.Println(k, top[k])
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
		case "/Андрей в дедлоке":
			// answer = answerStat(profile(playerUrl["Андрей"][0]))
			answer = answerStat(profile(playerUrl["Андрей"][0]))
			//space(answer)
		case "/Макс в дедлоке":
			answer = answerStat(profile(playerUrl["Макс"][0]))
			topWinrate()
		case "/Тимофей в дедлоке":
			answer = answerStat(profile(playerUrl["Тимофей"][0]))
		case "/Димасик в дедлоке":
			answer = answerStat(profile(playerUrl["Димасик"][0]))
		case "/Лёха в дедлоке":
			answer = answerStat(profile(playerUrl["Лёха"][0]))
		case "/Саша в дедлоке":
			answer = answerStat(profile(playerUrl["Саша"][0]))
		case "/deadlock":
			answer = "кек"
			topWinrate()
		default:
			answer = "You typed " + mess
		}
		return answer, nil
	})
}
