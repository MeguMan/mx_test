package store

type Store interface {
	Offer() OfferRepository
}