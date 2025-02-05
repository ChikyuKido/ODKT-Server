package helper

import (
	"bytes"
	"encoding/binary"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"odkt/server/db/entity"
	"odkt/server/db/repo"
	"odkt/server/util"
	"strings"
)

func ImportCardsToDB() {
	importCardChance()
	importCardBank()
	importCardSpecial()
	importCardStreet()
	importCardRailroad()
	importCardOther()
}

func importCardChance() {
	tree, err := toml.LoadFile("assets/dkt/cards/chance.toml")
	if err != nil {
		logrus.Fatalf("Failed to load file %v", err)
	}
	types := util.ConvertToStringArray(tree.Get("chance.types").([]interface{}))
	if err := repo.InsertSetting("card_chance", []byte(strings.Join(types, ","))); err != nil {
		logrus.Fatalf("Failed to add chance settings: %v", err)
	}
	entries := tree.Get("chances.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardChance := entity.CardChance{
			Type: getString(entryMap, "type"),
			Text: getString(entryMap, "text"),
		}
		var payload []byte
		if cardChance.Type == "pay" || cardChance.Type == "receive" || cardChance.Type == "land-tax" {
			payload = intToBytes(getInt(entryMap, "value"))
		} else if cardChance.Type == "move" {
			payload = []byte(getString(entryMap, "field"))
		} else if cardChance.Type == "house-repair" {
			payload = intToBytes(getInt(entryMap, "house"))
			payload = append(payload, intToBytes(getInt(entryMap, "hotel"))...)
		} else if cardChance.Type == "hold" {
			payload = []byte(getString(entryMap, "action"))
		}
		cardChance.Payload = payload
		if err := repo.InsertCardChance(cardChance); err != nil {
			logrus.Fatalf("Failed to add chance: %v", err)
		}
	}
}

func importCardBank() {
	tree, err := toml.LoadFile("assets/dkt/cards/bank.toml")
	if err != nil {
		logrus.Fatalf("Failed to load file %v", err)
	}
	types := util.ConvertToStringArray(tree.Get("bank.types").([]interface{}))
	if err := repo.InsertSetting("card_bank", []byte(strings.Join(types, ","))); err != nil {
		logrus.Fatalf("Failed to add bank settings: %v", err)
	}
	entries := tree.Get("banks.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardBank := entity.CardBank{
			Type: getString(entryMap, "type"),
			Text: getString(entryMap, "text"),
		}
		var payload []byte
		if cardBank.Type == "pay" || cardBank.Type == "receive" {
			payload = intToBytes(getInt(entryMap, "value"))
		} else if cardBank.Type == "move" {
			payload = []byte(getString(entryMap, "field"))
		}
		cardBank.Payload = payload
		if err := repo.InsertCardBank(cardBank); err != nil {
			logrus.Fatalf("Failed to add bank: %v", err)
		}
	}
}

func importCardSpecial() {
	tree, err := toml.LoadFile("assets/dkt/cards/special.toml")
	if err != nil {
		logrus.Fatalf("Failed to load file %v", err)
	}
	entryMap := tree.ToMap()["special"].(map[string]interface{})
	start := intToBytes(getInt(entryMap, "start"))
	multiplier := intToBytes(getInt(entryMap, "multiplier"))
	text := []byte(getString(entryMap, "text"))
	payload := start
	payload = append(payload, multiplier...)
	payload = append(payload, text...)
	if err := repo.InsertSetting("card_special", payload); err != nil {
		logrus.Fatalf("Failed to add special settings: %v", err)
	}
	entries := tree.Get("specials.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardSpecial := entity.CardSpecial{
			Name:      getString(entryMap, "name"),
			Price:     getInt(entryMap, "price"),
			PriceName: getString(entryMap, "price_name"),
		}

		if err := repo.InsertCardSpecial(cardSpecial); err != nil {
			logrus.Fatalf("Failed to add special: %v", err)
		}
	}
}

func importCardStreet() {
	tree, err := toml.LoadFile("assets/dkt/cards/street.toml")
	if err != nil {
		logrus.Fatalf("Failed to load file %v", err)
	}
	entries := tree.Get("streets.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardStreet := entity.CardStreet{
			City:       getString(entryMap, "city"),
			Name:       getString(entryMap, "name"),
			Rent:       util.ConvertIntArrayToString(util.ConvertToIntArray(entryMap["rent"].([]interface{})), ","),
			Price:      getInt(entryMap, "price"),
			HousePrice: getInt(entryMap, "price_house"),
			HotelPrice: getInt(entryMap, "price_hotel"),
		}
		if err := repo.InsertCardStreet(cardStreet); err != nil {
			logrus.Fatalf("Failed to add street: %v", err)
		}
	}
}

func importCardRailroad() {
	tree, err := toml.LoadFile("assets/dkt/cards/railroad.toml")
	if err != nil {
		logrus.Fatalf("Failed to load file %v", err)
	}
	entryMap := tree.ToMap()["railroad"].(map[string]interface{})
	start := intToBytes(getInt(entryMap, "start"))
	multiplier := intToBytes(getInt(entryMap, "multiplier"))
	text := []byte(getString(entryMap, "text"))
	payload := start
	payload = append(payload, multiplier...)
	payload = append(payload, text...)
	if err := repo.InsertSetting("card_railroad", payload); err != nil {
		logrus.Fatalf("Failed to add railroad settings: %v", err)
	}
	entries := tree.Get("railroads.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardRailroad := entity.CardRailroad{
			Name:  getString(entryMap, "name"),
			Price: getInt(entryMap, "price"),
		}
		if err := repo.InsertCardRailroad(cardRailroad); err != nil {
			logrus.Fatalf("Failed to add special: %v", err)
		}
	}
}
func importCardOther() {
	tree, err := toml.LoadFile("assets/dkt/cards/other-fields.toml")
	if err != nil {
		logrus.Fatalf("Failed to load file %v", err)
	}
	entries := tree.Get("other-fields.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardOther := entity.CardOther{
			Name: getString(entryMap, "name"),
		}
		var payload []byte
		if cardOther.Name == "vermögendsabgabe" {
			payload = intToBytes(getInt(entryMap, "percent"))
			payload = append(payload, intToBytes(getInt(entryMap, "maximum"))...)
		}
		cardOther.Payload = payload
		if err := repo.InsertCardOther(cardOther); err != nil {
			logrus.Fatalf("Failed to add special: %v", err)
		}
	}
}

func intToBytes(i int32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, i)
	if err != nil {
		logrus.Fatalf("binary.Write failed: %v", err)
	}
	return buf.Bytes()
}

func getUint(values map[string]interface{}, key string) uint {
	if v, ok := values[key].(int64); ok {
		return uint(v)
	}
	logrus.Fatalf("Failed to convert value %s to uint", key)
	return 0
}

func getInt(values map[string]interface{}, key string) int32 {
	if v, ok := values[key].(int64); ok {
		return int32(v)
	}
	logrus.Fatalf("Failed to convert value %s to int", key)
	return 0
}
func getString(values map[string]interface{}, key string) string {
	if v, ok := values[key].(string); ok {
		return v
	}
	logrus.Fatalf("Failed to convert value %s to string", key)
	return ""
}

func failOnError[T any](value T, err error) T {
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	return value
}
