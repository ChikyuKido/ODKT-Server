package repo

import (
	"odkt/server/db"
	"odkt/server/db/entity"
)

func InsertSetting(name string, payload []byte) error {
	if err := db.DB().Create(&entity.CardSettings{Name: name, Payload: payload}).Error; err != nil {
		return err
	}
	return nil
}
func InsertCardChance(card entity.CardChance) error {
	if err := db.DB().Create(&card).Error; err != nil {
		return err
	}
	return nil
}
func InsertCardBank(card entity.CardBank) error {
	if err := db.DB().Create(&card).Error; err != nil {
		return err
	}
	return nil
}
func InsertCardSpecial(card entity.CardSpecial) error {
	if err := db.DB().Create(&card).Error; err != nil {
		return err
	}
	return nil
}
func InsertCardStreet(card entity.CardStreet) error {
	if err := db.DB().Create(&card).Error; err != nil {
		return err
	}
	return nil
}

func InsertCardRailroad(card entity.CardRailroad) error {
	if err := db.DB().Create(&card).Error; err != nil {
		return err
	}
	return nil
}

func InsertCardOther(card entity.CardOther) error {
	if err := db.DB().Create(&card).Error; err != nil {
		return err
	}
	return nil
}
