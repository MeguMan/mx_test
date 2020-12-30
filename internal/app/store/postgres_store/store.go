package postgres_store

import (
	"github.com/jackc/pgx/v4"
)

type Store struct {
	conn *pgx.Conn
	OfferRepository *OfferRepository
}

func New(conn *pgx.Conn) *Store {
	return &Store{
		conn: conn,
	}
}

func (s *Store) Offer() *OfferRepository {
	if s.OfferRepository == nil {
		s.OfferRepository = &OfferRepository{
			store: s,
		}
	}

	return s.OfferRepository
}