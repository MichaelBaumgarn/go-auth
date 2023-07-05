package main

import "gorm.io/gorm"

type Grammar struct {
	db       *gorm.DB
	ID       uint   `gorm:"column:grammar_id" gorm:"primaryKey"`
	Language string `gorm:"column:language" json:"langauge" binding:"required"`
	Index    int    `gorm:"column:index" json:"index" binding:"required"`
	Word     string `gorm:"column:word" json:"word" binding:"required"`
	Complete string `gorm:"column:complete" json:"complete" binding:"required"`
}

type Tabler interface {
	TableName() string
}

func (Grammar) TableName() string {
	return "grammar"
}

func (m Grammar) Create(grammar Grammar) Grammar {
	m.db.Create(&grammar)
	return grammar
}

func (m Grammar) GetAll() []Grammar {
	var grammar []Grammar
	m.db.Find(&grammar)
	return grammar
}
