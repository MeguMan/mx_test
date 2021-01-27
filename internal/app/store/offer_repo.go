package store

import "github.com/MeguMan/mx_test/internal/app/model"

type OfferRepository interface {
	Create(o *model.Offer, rs *model.RowsStats) error
	GetByPattern(offerId, sellerId *int, pattern *string) ([]model.Offer, error)
	Delete(o *model.Offer, rs *model.RowsStats) error
	Exists(offerId, sellerId int) bool
	Update(o *model.Offer, rs *model.RowsStats) error
}