package main

import "gorm.io/gorm"

type UserGrammar struct {
	db        *gorm.DB
	ID        uint `gorm:"column:user_grammar_id" gorm:"primaryKey"`
	UserId    int  `gorm:"column:user_id" json:"userId" binding:"required"`
	GrammarId int  `gorm:"column:grammar_id" json:"grammarId" binding:"required"`
}

func (UserGrammar) TableName() string {
	return "user_grammar"
}

func (m UserGrammar) Create(grammar UserGrammar) UserGrammar {
	m.db.Create(&grammar)
	return grammar
}

func (m UserGrammar) GetAll() []UserGrammar {
	var grammar []UserGrammar
	m.db.Find(&grammar)
	return grammar
}
