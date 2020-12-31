package postgres_store

import (
	"context"
	"fmt"
	"github.com/MeguMan/mx_test/internal/app/model"
)

type OfferRepository struct {
	store *Store
}

func (r *OfferRepository) Create(o *model.Offer) error {
	//esli offerid = sellerid delete()
	_, err := r.store.conn.Exec(context.Background(), "INSERT INTO offers (offer_id, name, price, quantity, seller_id) VALUES ($1, $2, $3, $4, $5)",
		o.OfferId, o.Name, o.Price, o.Quantity, o.SellerId)
	return err
}

func (r *OfferRepository) GetAll(offerId, sellerId int, pattern string) error {
	rows, err := r.store.conn.Query(context.Background(), "SELECT * FROM offers WHERE offer_id = $1 AND seller_id = $2 AND name LIKE $3",
		offerId, sellerId, pattern + "%")
	for rows.Next() {
		fmt.Println(rows.Values())
	}

	return err
}

func (r *OfferRepository) Delete(o *model.Offer) error {
	_, err := r.store.conn.Exec(context.Background(), "DELETE FROM offers WHERE offer_id = $1", o.OfferId)
	return err
}