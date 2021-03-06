package repository

import (
	"context"
	. "github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
)

func (r *Repository) ConvertAmount(ctx context.Context, req *repository.ConvertRequest, rsp *repository.ConvertResponse) error {
	var rate *repository.ConvertResponse

	if err := r.GetConvertRate(ctx, req, rate); err != nil {
		return err
	}

	rsp.Amount = req.Amount / rate.Amount

	return nil
}

func (r *Repository) GetConvertRate(ctx context.Context, req *repository.ConvertRequest, rsp *repository.ConvertResponse) error {
	var rate *CurrencyRate

	err := r.Database.Collection(CollectionCurrencyRate).Find(
		bson.M{
			"currency_from": req.CurrencyFrom,
			"currency_to": req.CurrencyTo,
		}).One(&rate)

	if err != nil || rate == nil {
		log.Printf(QueryErrorMask, CollectionCurrencyRate, err.Error())
		return err
	}

	rsp.Amount = rate.GetRate()

	return nil
}
