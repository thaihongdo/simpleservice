package service

import (
	"errors"
	"simpleservice/cmd/entity-server/model"
	"time"
)

type BuyReq struct {
	WagerID     uint    `json:"wager_id"`
	BuyingPrice float64 `json:"buying_price"`
}

type BuyRes struct {
	ID          uint      `json:"id"`
	WagerID     uint      `json:"wager_id"`
	BuyingPrice float64   `json:"buying_price"`
	BoughtAt    time.Time `json:"bought_at"`
}

func (obj *BuyReq) Buy() (interface{}, error) {
	record := model.Buy{
		WagerID:     obj.WagerID,
		BuyingPrice: obj.BuyingPrice,
	}

	resObj, err := record.Add()
	if err != nil {
		return nil, errors.New("create record fail")
	}
	//after create success
	//we need update wager
	// (&model.Wager{
	// 	CurrentSellingPrice: 1,
	// 	PercentageSold:      1,
	// 	AmountSold:          1,
	// }).UpdateAfterBuy(obj.WagerID)

	return toBuyResponse(resObj), nil
}

func toBuyResponse(model *model.Buy) interface{} {
	return &BuyRes{
		ID:          model.ID,
		WagerID:     model.WagerID,
		BuyingPrice: model.BuyingPrice,
		BoughtAt:    model.CreatedAt,
	}
}
