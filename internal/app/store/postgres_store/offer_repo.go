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
	_, err := r.store.conn.Exec(context.Background(), "INSERT INTO offers (offer_id, name, price, quantity, seller_id) VALUES ($1, $2, $3, $4, $5)",
		o.OfferId, o.Name, o.Price, o.Quantity, o.SellerId)
	return err
}

func (r *OfferRepository) GetAll() error {
	rows, err := r.store.conn.Query(context.Background(), "SELECT * FROM offers")
	fmt.Println("ROWS -> ", rows)
	for rows.Next() {
		fmt.Println(rows.Values())
	}
	return err
}