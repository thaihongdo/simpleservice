package model

import (
	"gorm.io/gorm"
)

/*
Define for shape:
Square: Length
Rectangle: Length, Length1
Triangle : Length, Length1, Length2
Diamond: Length, Length1
*/

type Buy struct {
	gorm.Model

	WagerID     uint    `json:"wager_id"`
	BuyingPrice float64 `json:"buying_price"`
}

func (obj *Buy) Add() (*Buy, error) {
	if err := db.Create(obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}
