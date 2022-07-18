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

type Wager struct {
	gorm.Model

	TotalWagerValue     int     `json:"total_wager_value"`
	Odds                int     `json:"odds"`
	SellingPercentage   int     `json:"selling_percentage"`
	SellingPrice        float64 `json:"selling_price"`
	CurrentSellingPrice float64 `json:"current_selling_price"`
	PercentageSold      float64 `json:"percentage_sold"`
	AmountSold          float64 `json:"amount_sold"`
}

func (obj *Wager) Get(ID uint) (*Wager, error) {
	err := db.
		Where("id = ?", obj.ID).
		First(obj).Error
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (obj *Wager) GetWager(pageNum int, pageSize int) ([]*Wager, error) {
	var list []*Wager
	err := db.Offset(pageNum).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return list, nil
}

func (obj *Wager) Add() (*Wager, error) {
	if err := db.Create(obj).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

func (obj *Wager) UpdateAfterBuy(id uint) {
	var tmpObj Wager
	err := db.Where("id = ?", id).First(&tmpObj).Error
	if err != nil {
		return
	}
	tmpObj.AmountSold = obj.AmountSold
	tmpObj.PercentageSold = obj.PercentageSold
	tmpObj.CurrentSellingPrice = obj.CurrentSellingPrice

	db.Save(&tmpObj)
}
