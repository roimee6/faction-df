package util

var Ranks = map[string]map[string]string{
	"joueur": {
		"gamertag": "§7[{fac}] {name}",
		"chat":     "§7[{fac}] {name}: {msg}",
	},
	"guide": {
		"gamertag": "§a[{fac}] {name}",
		"chat":     "§a[{fac}][Guide] {name}: {msg}",
	},
	"fondateur": {
		"gamertag": "§4[{fac}] {name}",
		"chat":     "§4[{fac}][Fonda] {name}: {msg}",
	},
}

var RanksList = []string{
	"joueur",
	"guide",
	"fondateur",
}

var WelcomeMessages = []string{
	"Bienvenue à toi @{player} ! Passe du bon temps sur le serveur !",
	"Hey @{player} ! Je te souhaite la bienvenue parmi nous",
	"Salut @{player} nous sommes content de taccueillir aujourdhui et jespère que tu te plairas ici",
	"Heureux de te rencontrer @{player} fais comme chez toi",
	"Bienvenue sur le serveur @{player} ! Tu nous amène des pizzas ?",
}

var Zones = map[string]string{
	"spawn": "99:0:100:-100:255:-99",
}
