package postgres_store

import (
	"context"
	"github.com/MeguMan/mx_test/internal/app/model"
)

type OfferRepository struct {
	store *Store
}

func (r *OfferRepository) Create(o *model.Offer) error {
	exists := r.Exists(o.OfferId, o.SellerId)
	if exists {
		err := r.Update(o)
		return err
	}
	
	_, err := r.store.conn.Exec(context.Background(), "INSERT INTO offers (offer_id, name, price, quantity, seller_id) VALUES ($1, $2, $3, $4, $5)",
		o.OfferId, o.Name, o.Price, o.Quantity, o.SellerId)
	return err
}

func (r *OfferRepository) GetByPattern(offerId, sellerId int, pattern string) ([]model.Offer, error) {
	rows, err := r.store.conn.Query(context.Background(), "SELECT * FROM offers WHERE offer_id = $1 AND seller_id = $2 AND name LIKE $3",
		offerId, sellerId, pattern + "%")

	var oo []model.Offer
	for rows.Next() {
		o := model.Offer{}
		err := rows.Scan(&o.OfferId, &o.Name, &o.Price, &o.Quantity, &o.SellerId)
		o.Available = true
		if err != nil {
			return nil, err
		}
		oo = append(oo, o)
	}

	return oo, err
}

func (r *OfferRepository) Delete(o *model.Offer) error {
	_, err := r.store.conn.Exec(context.Background(), "DELETE FROM offers WHERE offer_id = $1", o.OfferId)
	return err
}

func (r *OfferRepository) Exists(offerId, sellerId int) bool {
	var exists bool
	_ = r.store.conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM offers WHERE offer_id = $1 AND seller_id = $2)",
		offerId, sellerId).Scan(&exists)
	return exists
}

func (r *OfferRepository) Update(o *model.Offer) error {
	_, err := r.store.conn.Exec(context.Background(), "UPDATE offers SET name = $1, price = $2, quantity = $3 WHERE seller_id = $4 and offer_id = $5",
		o.Name, o.Price, o.Quantity, o.SellerId, o.OfferId)
	return err
}