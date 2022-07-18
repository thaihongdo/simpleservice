package service

import (
	"errors"
	"math"
	"simpleservice/cmd/entity-server/model"
	"time"
)

type WagerReq struct {
	ID                uint    `json:"id"`
	TotalWagerValue   int     `json:"total_wager_value"`
	Odds              int     `json:"odds"`
	SellingPercentage int     `json:"selling_percentage" valid:"range(0|100)~Selling Percentage range 0 to 100"`
	SellingPrice      float64 `json:"selling_price"`
	PageNum           int
	PageSize          int
}

type WagerRes struct {
	ID                  uint      `json:"id"`
	TotalWagerValue     int       `json:"total_wager_value"`
	Odds                int       `json:"odds"`
	SellingPercentage   int       `json:"selling_percentage"`
	SellingPrice        float64   `json:"selling_price"`
	CurrentSellingPrice float64   `json:"current_selling_price"`
	PercentageSold      float64   `json:"percentage_sold"`
	AmountSold          float64   `json:"amount_sold"`
	PlacedAt            time.Time `json:"placed_at"`
}

func (obj *WagerReq) Get() (interface{}, error) {
	record := model.Wager{}
	resObj, err := record.Get(obj.ID)
	if err != nil {
		return nil, errors.New("record not found")
	}

	return toResponse(resObj), nil
}

func (obj *WagerReq) GetWager() ([]*WagerRes, error) {
	var res []*WagerRes
	record := model.Wager{}
	list, err := record.GetWager(obj.PageNum, obj.PageSize)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		objRes := toResponse(item)
		res = append(res, objRes)
	}
	return res, nil
}

func (obj *WagerReq) Add() (interface{}, error) {
	sellingPrice := math.Round(obj.SellingPrice*100) / 100
	numberCompare := obj.TotalWagerValue * (obj.SellingPercentage / 100)

	if sellingPrice < float64(numberCompare) {
		return nil, errors.New("selling price invalid")
	}
	record := model.Wager{
		TotalWagerValue:   obj.TotalWagerValue,
		Odds:              obj.Odds,
		SellingPrice:      sellingPrice,
		SellingPercentage: obj.SellingPercentage,
	}

	resObj, err := record.Add()
	if err != nil {
		return nil, errors.New("create record fail")
	}

	return toResponse(resObj), nil
}

func toResponse(model *model.Wager) *WagerRes {
	return &WagerRes{
		ID:                  model.ID,
		TotalWagerValue:     model.TotalWagerValue,
		Odds:                model.Odds,
		SellingPercentage:   model.SellingPercentage,
		SellingPrice:        model.SellingPrice,
		CurrentSellingPrice: model.CurrentSellingPrice,
		PercentageSold:      model.PercentageSold,
		AmountSold:          model.AmountSold,
		PlacedAt:            model.CreatedAt,
	}
}
