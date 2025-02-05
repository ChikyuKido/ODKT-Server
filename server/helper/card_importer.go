package helper

import (
	"bytes"
	"encoding/binary"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"log"
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
}

func importCardChance() {
	tree, err := toml.LoadFile("assets/cards/chance.toml")
	if err != nil {
		log.Fatal(err)
	}
	types := util.ConvertToStringArray(tree.Get("chance.typen").([]interface{}))
	if err := repo.InsertSetting("card_chance", []byte(strings.Join(types, ","))); err != nil {
		logrus.Fatalf("Failed to add chance settings: %v", err)
	}
	entries := tree.Get("chance.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardChance := entity.CardChance{
			Type: getString(entryMap, "typ"),
			Text: getString(entryMap, "text"),
		}
		var payload []byte
		if cardChance.Type == "zahle" || cardChance.Type == "erhalte" || cardChance.Type == "grundsteuer" {
			payload = intToBytes(getInt(entryMap, "wert"))
		} else if cardChance.Type == "gehe" {
			payload = []byte(getString(entryMap, "feld"))
		} else if cardChance.Type == "haus-reperatur" {
			payload = intToBytes(getInt(entryMap, "haus"))
			payload = append(payload, intToBytes(getInt(entryMap, "hotel"))...)
		} else if cardChance.Type == "hebe" {
			payload = []byte(getString(entryMap, "aktion"))
		}
		cardChance.Payload = payload
		if err := repo.InsertCardChance(cardChance); err != nil {
			logrus.Fatalf("Failed to add chance: %v", err)
		}
	}
}

func importCardBank() {
	tree, err := toml.LoadFile("assets/cards/sparkassa.toml")
	if err != nil {
		log.Fatal(err)
	}
	types := util.ConvertToStringArray(tree.Get("sparkassa.typen").([]interface{}))
	if err := repo.InsertSetting("card_bank", []byte(strings.Join(types, ","))); err != nil {
		logrus.Fatalf("Failed to add bank settings: %v", err)
	}
	entries := tree.Get("sparkassa.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardBank := entity.CardBank{
			Type: getString(entryMap, "typ"),
			Text: getString(entryMap, "text"),
		}
		var payload []byte
		if cardBank.Type == "zahle" || cardBank.Type == "erhalte" {
			payload = intToBytes(getInt(entryMap, "wert"))
		} else if cardBank.Type == "gehe" {
			payload = []byte(getString(entryMap, "feld"))
		}
		cardBank.Payload = payload
		if err := repo.InsertCardBank(cardBank); err != nil {
			logrus.Fatalf("Failed to add bank: %v", err)
		}
	}
}

func importCardSpecial() {
	tree, err := toml.LoadFile("assets/cards/spezial.toml")
	if err != nil {
		log.Fatal(err)
	}
	entryMap := tree.ToMap()["spezial"].(map[string]interface{})
	start := intToBytes(getInt(entryMap, "start"))
	multiplier := intToBytes(getInt(entryMap, "multiplier"))
	text := []byte(getString(entryMap, "text"))
	payload := start
	payload = append(payload, multiplier...)
	payload = append(payload, text...)
	if err := repo.InsertSetting("card_special", payload); err != nil {
		logrus.Fatalf("Failed to add special settings: %v", err)
	}
	entries := tree.Get("spezial.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardSpecial := entity.CardSpecial{
			Name:      getString(entryMap, "name"),
			Price:     getInt(entryMap, "preis"),
			PriceName: getString(entryMap, "preis_name"),
		}
		if err := repo.InsertCardSpecial(cardSpecial); err != nil {
			logrus.Fatalf("Failed to add special: %v", err)
		}
	}
}

func importCardStreet() {
	tree, err := toml.LoadFile("assets/cards/strasse.toml")
	if err != nil {
		log.Fatal(err)
	}
	entries := tree.Get("strassen.entry").([]*toml.Tree)
	for _, entry := range entries {
		entryMap := entry.ToMap()
		cardStreet := entity.CardStreet{
			City:       getString(entryMap, "stadt"),
			Name:       getString(entryMap, "name"),
			Rent:       util.ConvertIntArrayToString(util.ConvertToIntArray(entryMap["miete"].([]interface{})), ","),
			Price:      getInt(entryMap, "preis"),
			HousePrice: getInt(entryMap, "preis_haus"),
			HotelPrice: getInt(entryMap, "preis_hotel"),
		}
		if err := repo.InsertCardStreet(cardStreet); err != nil {
			logrus.Fatalf("Failed to add street: %v", err)
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
