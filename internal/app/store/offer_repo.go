package store

import "github.com/MeguMan/mx_test/internal/app/model"

type OfferRepository interface {
	Create(o *model.Offer) error
}