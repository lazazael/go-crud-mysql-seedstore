package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lazazael/go-crud-mysql-seedstore/pkg/config"
)

var db *gorm.DB

type Seed struct {
	gorm.Model
	SpeciesName string `gorm:"" json:"species"`
	VarietyName string `json:"variety"`
	//Classification string `json:"classification"`
	PacketContent int `json:"packet-content"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Seed{})
}

func (s *Seed) CreateSeed() *Seed {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllSeeds() []Seed {
	var Seeds []Seed
	db.Find(&Seeds)
	return Seeds
}

func GetSeedById(Id int64) (*Seed, *gorm.DB) {
	var getSeed Seed
	db := db.Where("ID=?", Id).Find(&getSeed)
	return &getSeed, db
}

func DeleteSeed(Id int64) Seed {
	var seed Seed
	db.Where("Id=?", Id).Delete(seed)
	return seed
}

//UpdateSeed is get+delete+create
