package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var language = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Русский", "ru"),
		tgbotapi.NewInlineKeyboardButtonData("English", "en"),
	),
)

var kbru = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Аниме"),
		tgbotapi.NewKeyboardButton("Погода"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Язык"),
	),
)
var kben = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Anime"),
		tgbotapi.NewKeyboardButton("Weather"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Language"),
	),
)
var animekben = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1/2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("I was kicked out of the guild of heroes, season 2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Lunar travel will lead to a new world, season 2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Leveling Up Alone"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Monologue of a pharmacist"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Seeing off Friren on her last journey"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Further➡️"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Back"),
	),
)
var animekbru = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1/2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Меня выгнали из гильдии героев, 2 сезон"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Лунное путешествие приведёт к новому миру, 2 сезон"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Поднятие уровня в одиночку"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Монолог фармацевта"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Провожающая в последний путь Фрирен"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Дальше➡️"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)
var animekb1ru = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("2/2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Добро пожаловать в класс превосходства, 3 сезон"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Низкоуровневый персонаж Томодзаки, 2 сезон"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Другой мир не может противостоять силе мгновенной смерти"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Невероятный танк проходит подземелья"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Злодейка наслаждается своей седьмой жизнью в качестве свободолюбивой невесты во вражеской стране"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⬅️Назад"),
	),
)
var animekb1en = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("2/2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Welcome to Superiority Class Season 3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Low level character Tomozaki, season 2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Another world cannot withstand the power of instant death"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Incredible tank goes through dungeons"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("The villainess is enjoying her seventh life as a freedom-loving bride in an enemy country"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("⬅Back"),
	),
)
var nkbru = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Москва", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Санкт-Петербург", "2"),
		tgbotapi.NewInlineKeyboardButtonData("Новосибирск", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Екатеринбург", "4"),
		tgbotapi.NewInlineKeyboardButtonData("Казань", "5"),
		tgbotapi.NewInlineKeyboardButtonData("Нижний Новгород", "6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Красноярск", "7"),
		tgbotapi.NewInlineKeyboardButtonData("Челябинск", "8"),
		tgbotapi.NewInlineKeyboardButtonData("Самара", "9"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Уфа", "10"),
		tgbotapi.NewInlineKeyboardButtonData("Ростов-на-Дону", "11"),
		tgbotapi.NewInlineKeyboardButtonData("Краснодар", "12"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Омск", "13"),
		tgbotapi.NewInlineKeyboardButtonData("Воронеж", "14"),
		tgbotapi.NewInlineKeyboardButtonData("Пермь", "15"),
	),
)
var nkben = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Moscow", "1"),
		tgbotapi.NewInlineKeyboardButtonData("Saint-Petersburg", "2"),
		tgbotapi.NewInlineKeyboardButtonData("Novosibirsk", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ekaterinburg", "4"),
		tgbotapi.NewInlineKeyboardButtonData("Kazan", "5"),
		tgbotapi.NewInlineKeyboardButtonData("Lower-Novgorod", "6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Krasnoyarsk", "7"),
		tgbotapi.NewInlineKeyboardButtonData("Chelyabinsk", "8"),
		tgbotapi.NewInlineKeyboardButtonData("Samara", "9"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Ufa", "10"),
		tgbotapi.NewInlineKeyboardButtonData("Rostov-on-Don", "11"),
		tgbotapi.NewInlineKeyboardButtonData("Krasnodar", "12"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Omsk", "13"),
		tgbotapi.NewInlineKeyboardButtonData("Voronezh", "14"),
		tgbotapi.NewInlineKeyboardButtonData("Permian", "15"),
	),
)

var answers = []string{
	"Привет",
	"Hi",
	"Приветсвую",
	"Каничива",
	"Здравствуйте",
	"Как ваши ничего?",
	"Салам",
	"Саламалейкум",
}
var answers2 = []string{
	"Нормально,у вас как?",
	"Бывало и получше)",
	"Да в целом, все норм",
	"Сойдет",
	"Я рад что вы спросили)",
	"Да все хорошо,вы как?",
	"До этого момента все было отлично, а теперь стало супер)",
}

var MetioAndAnime = map[string]gis{
	"погода всеволожск":      {"Погода на сегодня, в городе Всеволожск", "https://www.gismeteo.ru/weather-vsevolozhsk-11425/", "Посмотреть вы можете тут 👉"},
	"погода москва":          {"Погода на сегодня, в городе Москва", "https://www.gismeteo.ru/weather-moscow-4368/", "Посмотреть вы можете тут 👉"},
	"погода санкт-питербург": {"Погода на сегодня, в городе Санкт-Петербург", "https://www.gismeteo.ru/weather-sankt-peterburg-4079/", "Посмотреть вы можете тут 👉"},
	"погода новосибирск":     {"Погода на сегодня, в городе Новосибирск", "https://www.gismeteo.ru/weather-novosibirsk-4690/", "Посмотреть вы можете тут 👉"},
	"погода екатеринбург":    {"Погода на сегодня, в городе Екатеринбург", "https://www.gismeteo.ru/weather-yekaterinburg-4517/", "Посмотреть вы можете тут 👉"},
	"погода казань":          {"Погода на сегодня, в городе Казань", "https://www.gismeteo.ru/weather-kazan-4364/", "Посмотреть вы можете тут 👉"},
	"погода нижний новгород": {"Погода на сегодня, в городе Нижний Новгород", "https://www.gismeteo.ru/weather-nizhny-novgorod-4355/", "Посмотреть вы можете тут 👉"},
	"погода красноярск":      {"Погода на сегодня, в городе Красноярск", "https://www.gismeteo.ru/weather-krasnoyarsk-4674/", "Посмотреть вы можете тут 👉"},
	"погода челябинск":       {"Погода на сегодня, в городе Челябинск", "https://www.gismeteo.ru/weather-chelyabinsk-4565/", "Посмотреть вы можете тут 👉"},
	"погода самара":          {"Погода на сегодня, в городе Самара", "https://www.gismeteo.ru/weather-samara-4618/", "Посмотреть вы можете тут 👉"},
	"погода уфа":             {"Погода на сегодня, в городе Уфа", "https://www.gismeteo.ru/weather-ufa-4588/", "Посмотреть вы можете тут 👉"},
	"погода ростов на дону":  {"Погода на сегодня, в городе Ростов-На-Дону", "https://www.gismeteo.ru/weather-rostov-na-donu-5110/", "Посмотреть вы можете тут 👉"},
	"погода краснодар":       {"Погода на сегодня, в городе Краснодар", "https://www.gismeteo.ru/weather-krasnodar-5136/", "Посмотреть вы можете тут 👉"},
	"погода омск":            {"Погода на сегодня, в городе Омск", "https://www.gismeteo.ru/weather-omsk-4578/", "Посмотреть вы можете тут 👉"},
	"погода воронеж":         {"Погода на сегодня, в городе Воронеж", "https://www.gismeteo.ru/weather-voronezh-5026/", "Посмотреть вы можете тут 👉"},
	"погода пермь":           {"Погода на сегодня, в городе Пермь", "https://www.gismeteo.ru/weather-perm-4476/", "Посмотреть вы можете тут 👉"},
	///.
	"аниме на сегодня":                                         {"Онгоинг аниме на сегодня", "https://shikimori.one/ongoings", "Посмотреть вы можете тут 👉"},
	"меня выгнали из гильдии героев, 2 сезон":                  {"Aниме на сегодня", "https://animego.org/anime/menya-vygnali-iz-gildii-geroev-potomu-chto-ya-byl-plohim-kompanonom-poetomu-ya-reshil-nespeshno-zhit-v-glushi-2-2470", "Посмотреть вы можете тут 👉"},
	"лунное путешествие приведёт к новому миру, 2 сезон":       {"Aниме на сегодня", "https://animego.org/anime/lunnoe-puteshestvie-privedet-k-novomu-miru-2-2463", "Посмотреть вы можете тут 👉"},
	"поднятие уровня в одиночку":                               {"Aниме на сегодня", "https://animego.org/anime/podnyatie-urovnya-v-odinochku-2477", "Посмотреть вы можете тут 👉"},
	"Монолог фармацевта":                                       {"Aниме на сегодня", "https://animego.org/anime/monolog-farmacevta-2422", "Посмотреть вы можете тут 👉"},
	"провожающая в последний путь фрирен":                      {"Aниме на сегодня", "https://animego.org/anime/provozhayuschaya-v-posledniy-put-friren-2430", "Посмотреть вы можете тут 👉"},
	"добро пожаловать в класс превосходства, 3 сезон":          {"Aниме на сегодня", "https://animego.org/anime/dobro-pozhalovat-v-klass-prevoshodstva-3-2465", "Посмотреть вы можете тут 👉"},
	"низкоуровневый персонаж томодзаки, 2 сезон":               {"Aниме на сегодня", "https://animego.org/anime/nizkourovnevyy-personazh-tomodzaki-vtoraya-stadiya-2464", "Посмотреть вы можете тут 👉"},
	"другой мир не может противостоять силе мгновенной смерти": {"Aниме на сегодня", "https://animego.org/anime/drugoy-mir-ne-mozhet-protivostoyat-sile-mgnovennoy-smerti-2476", "Посмотреть вы можете тут 👉"},
	"невероятный танк проходит подземелья":                     {"Aниме на сегодня", "https://animego.org/anime/neveroyatnyy-tank-prohodit-podzemelya-tank-obladayuschiy-redkim-navykom-vynoslivosti-9999-byl-izgnan-iz-geroyskoy-gruppy-2500", "Посмотреть вы можете тут 👉"},
	"злодейка наслаждается своей седьмой жизнью в качестве свободолюбивой невесты во вражеской стране": {"Aниме на сегодня", "https://animego.org/anime/zlodeyka-naslazhdaetsya-svoey-sedmoy-zhiznyu-v-kachestve-svobodolyubivoy-nevesty-vo-vrazheskoy-strane-2496", "Посмотреть вы можете тут 👉"},
	///.
	"i was kicked out of the guild of heroes, season 2": {"Anime for today", "https://animego.org/anime/menya-vygnali-iz-gildii-geroev-potomu-chto-ya-byl-plohim-kompanonom-poetomu-ya-reshil-nespeshno-zhit-v-glushi-2-2470", "You can watch it here 👉"},
	"lunar travel will lead to a new world, season 2":   {"Anime for today", "https://animego.org/anime/lunnoe-puteshestvie-privedet-k-novomu-miru-2-2463", "You can watch it here 👉"},
	"leveling up alone":                                         {"Anime for today", "https://animego.org/anime/podnyatie-urovnya-v-odinochku-2477", "You can watch it here 👉"},
	"monologue of a pharmacist":                                 {"Anime for today", "https://animego.org/anime/monolog-farmacevta-2422", "You can watch it here 👉"},
	"seeing off friren on her last journey":                     {"Anime for today", "https://animego.org/anime/provozhayuschaya-v-posledniy-put-friren-2430", "You can watch it here 👉"},
	"welcome to superiority class season 3":                     {"Anime for today", "https://animego.org/anime/dobro-pozhalovat-v-klass-prevoshodstva-3-2465", "You can watch it here 👉"},
	"low level character tomozaki, season 2":                    {"Anime for today", "https://animego.org/anime/nizkourovnevyy-personazh-tomodzaki-vtoraya-stadiya-2464", "You can watch it here 👉"},
	"another world cannot withstand the power of instant death": {"Anime for today", "https://animego.org/anime/drugoy-mir-ne-mozhet-protivostoyat-sile-mgnovennoy-smerti-2476", "You can watch it here 👉"},
	"incredible tank goes through dungeons":                     {"Anime for today", "https://animego.org/anime/neveroyatnyy-tank-prohodit-podzemelya-tank-obladayuschiy-redkim-navykom-vynoslivosti-9999-byl-izgnan-iz-geroyskoy-gruppy-2500", "You can watch it here 👉"},
	"the villainess is enjoying her seventh life as a freedom-loving bride in an enemy country": {"Anime for today", "https://animego.org/anime/zlodeyka-naslazhdaetsya-svoey-sedmoy-zhiznyu-v-kachestve-svobodolyubivoy-nevesty-vo-vrazheskoy-strane-2496", "You can watch it here 👉"},
}

var MetioData = map[string]gisqq{
	"1":  {"Погода на сегодня, в городе Москва", "https://www.gismeteo.ru/weather-moscow-4368/"},
	"2":  {"Погода на сегодня, в городе Санкт-Петербург", "https://www.gismeteo.ru/weather-sankt-peterburg-4079/"},
	"3":  {"Погода на сегодня, в городе Новосибирск", "https://www.gismeteo.ru/weather-novosibirsk-4690/"},
	"4":  {"Погода на сегодня, в городе Екатеринбург", "https://www.gismeteo.ru/weather-yekaterinburg-4517/"},
	"5":  {"Погода на сегодня, в городе Казань", "https://www.gismeteo.ru/weather-kazan-4364/"},
	"6":  {"Погода на сегодня, в городе Нижний Новгород", "https://www.gismeteo.ru/weather-nizhny-novgorod-4355/"},
	"7":  {"Погода на сегодня, в городе Красноярск", "https://www.gismeteo.ru/weather-krasnoyarsk-4674/"},
	"8":  {"Погода на сегодня, в городе Челябинск", "https://www.gismeteo.ru/weather-chelyabinsk-4565/"},
	"9":  {"Погода на сегодня, в городе Самара", "https://www.gismeteo.ru/weather-samara-4618/"},
	"10": {"Погода на сегодня, в городе Уфа", "https://www.gismeteo.ru/weather-ufa-4588/"},
	"11": {"Погода на сегодня, в городе Ростов-На-Дону", "https://www.gismeteo.ru/weather-rostov-na-donu-5110/"},
	"12": {"Погода на сегодня, в городе Краснодар", "https://www.gismeteo.ru/weather-krasnodar-5136/"},
	"13": {"Погода на сегодня, в городе Омск", "https://www.gismeteo.ru/weather-omsk-4578/"},
	"14": {"Погода на сегодня, в городе Воронеж", "https://www.gismeteo.ru/weather-voronezh-5026/"},
	"15": {"Погода на сегодня, в городе Пермь", "https://www.gismeteo.ru/weather-perm-4476/"},
}
var KeyBut = map[string]kbutt{
	"ru":        {"Вы выбрали Русский язык!", kbru},
	"en":        {"You chose English!", kben},
	"аниме":     {"Выберете из списка", animekbru},
	"anime":     {"Select" + " from the list", animekben},
	"дальше➡️":  {"Дальше➡️", animekb1ru},
	"further➡️": {"further➡️", animekb1en},
	"⬅️назад":   {"⬅️Назад", animekbru},
	"⬅back":     {"⬅Back", animekben},
	"back":      {"Back...", kben},
	"назад":     {"Назад...", kbru},
}
var KeyBytIrl = map[string]kbuttirl{
	"язык":     {"Смена языка", language},
	"language": {"Language", language},
	"погода":   {"Выберете город", nkbru},
	"weather":  {"Select a city", nkben},
	"/start":   {"Language", language},
}
var helo = map[string]helloy{
	"привет":                {"Приветсвую!"},
	"привет!":               {"Приветсвую!"},
	"здрасте":               {"Приветсвую!"},
	"приветствую":           {"Приветсвую!"},
	"здравствуйте":          {"Приветсвую!"},
	"добрый день":           {"Приветсвую!"},
	"добрый вечер":          {"Приветсвую!"},
	"доброе утро":           {"Приветсвую!"},
	"доброго времени суток": {"Приветсвую!"},
	"прив":                  {"Приветсвую!"},
	"хай":                   {"Приветсвую!"},
	"каничива":              {"Приветсвую!"},
	"салам":                 {"Приветсвую!"},
	"саламалейкум":          {"Приветсвую!"},
}
