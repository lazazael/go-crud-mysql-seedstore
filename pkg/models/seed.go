package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lazazael/go-crud-mysql-seedstore/config"
)

var db *gorm.DB

type Seed struct {
	gorm.model
	SpeciesName    string `gorm:""json:"species"`
	VarietyName    string `json:"variety"`
	Classification string `json:"classification"`
	PacketContent  int    `json:"packet-content"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Seed{})
}
