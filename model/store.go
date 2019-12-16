package model

import (
	"github.com/jinzhu/gorm"
)

type Store struct {
	gorm.Model
	StoreID     string `gorm:"primary_key"`
	Storename   string `gorm:"varchar(255);"`
	Address     string `gorm:"varchar(255);"`
	Genre       string `gorm:"varchar(255);"`
	Tel         string `gorm:"varchar(255);"`
	Information string `gorm:"varchar(255);"`
	Tables      string `gorm:"varchar(255);"`
}

func GetAllStores(db *gorm.DB) ([]Store, error) {
	stores := make([]Store, 10)
	if err := db.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func GetStoreByID(db *gorm.DB, StoreID uint) (*Store, error) {
	store := new(Store)
	if err := db.Where("StoreID = ?", StoreID).First(&store).Error; err != nil {
		return nil, err
	}
	return store, nil
}

func CreateStore(db *gorm.DB, store *Store) error {
	return db.Create(&store).Error
}

func DeleteStoreByID(db *gorm.DB, StoreID uint) error {
	return db.Where("StoreID = ?", StoreID).Delete(Store{}).Error
}
