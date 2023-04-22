package util

import (
	"encoding/json"
	"os"
)

var (
	Factions = make(map[string]*Faction)
)

func LoadCache() {
	loadFactions()
}

func loadFactions() {
	if _, err := os.Stat("asset/factions.json"); os.IsNotExist(err) {
		file, err := os.Create("asset/factions.json")
		if err != nil {
			panic(err)
		}

		_, err = file.WriteString("{}")
		if err != nil {
			panic(err)
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}

	data, err := os.ReadFile("asset/factions.json")
	if err != nil {
		panic(err)
	}

	var factions map[string]Faction

	err = json.Unmarshal(data, &factions)
	if err != nil {
		panic(err)
	}

	for index, element := range factions {
		Factions[index] = &element
	}
}
