package handler

import (
	"encoding/json"
	"github.com/roimee6/Faction/server/session"
	"github.com/roimee6/Faction/server/util"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	Factions = make(map[string]*util.Faction)
	Players  = make(map[string]map[string]interface{})

	AlreadyWished []string
	LastJoin      string
)

func LoadCache() {
	files, err := os.ReadDir("asset/data/players/")
	if err != nil {
		log.Fatal(err)
	}

	Players["xuid"] = make(map[string]interface{})
	Players["kill"] = make(map[string]interface{})
	Players["death"] = make(map[string]interface{})
	Players["money"] = make(map[string]interface{})
	Players["addresses"] = make(map[string]interface{})
	Players["uuids"] = make(map[string]interface{})
	Players["selfsignedids"] = make(map[string]interface{})
	Players["deviceids"] = make(map[string]interface{})

	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()

			extension := filepath.Ext(filename)
			nameWithoutExt := filename[0 : len(filename)-len(extension)]

			data := session.ParseData(nameWithoutExt)
			name := strings.ToLower(data.Name)

			PushData(nameWithoutExt, name, data)
		}
	}

	loadFactions()
}

func PushData(xuid string, name string, data session.Data) {
	name = strings.ToLower(name)
	Players["xuid"][name] = xuid

	v := reflect.ValueOf(data)

	for i := 0; i < v.NumField(); i++ {
		field := strings.ToLower(v.Type().Field(i).Name)
		value := v.Field(i).Interface()

		if _, ok := Players[field]; ok {
			Players[field][name] = value
		}
	}
}

func SaveCache() {
	file, err := os.Create("asset/factions.json")
	if err != nil {
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	jsonData, err := json.Marshal(Factions)
	if err != nil {
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return
	}
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

	var factions map[string]util.Faction

	err = json.Unmarshal(data, &factions)
	if err != nil {
		panic(err)
	}

	for index, element := range factions {
		Factions[index] = &element
	}
}
