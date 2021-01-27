package postgres_store

import (
	"context"
	"fmt"
	"github.com/MeguMan/mx_test/internal/app/model"
	"github.com/jackc/pgx/v4"
)

type OfferRepository struct {
	store *Store
}

func (r *OfferRepository) Create(o *model.Offer, rs *model.RowsStats) error {
	fmt.Println("create")
	_, err := r.store.conn.Exec(context.Background(), "INSERT INTO offers (offer_id, name, price, quantity, seller_id) VALUES ($1, $2, $3, $4, $5)",
		o.OfferId, o.Name, o.Price, o.Quantity, o.SellerId)
	fmt.Println(o.OfferId, o.Name, o.Price, o.Quantity, o.SellerId)
	rs.CreatedRows += 1
	return err
}

func (r *OfferRepository) GetByPattern(offerId, sellerId *int, pattern *string) ([]model.Offer, error) {
	var err error
	var rows pgx.Rows

	if pattern == nil {
		if offerId == nil && sellerId == nil {
			rows, err = r.store.conn.Query(context.Background(), "SELECT * FROM offers")
			if err != nil {
				return nil, err
			}
		} else if offerId == nil && sellerId != nil {
			rows, err = r.store.conn.Query(context.Background(), "SELECT * FROM offers WHERE seller_id = $1",
				sellerId)
			if err != nil {
				return nil, err
			}
		} else if offerId != nil && sellerId == nil {
			rows, err = r.store.conn.Query(context.Background(), "SELECT * FROM offers WHERE offer_id = $1 AND",
				offerId)
			if err != nil {
				return nil, err
			}
		} else {
			rows, err = r.store.conn.Query(context.Background(), "SELECT * FROM offers WHERE offer_id = $1 AND seller_id = $2",
				offerId, sellerId)
			if err != nil {
				return nil, err
			}
		}
	} else {
		if offerId == nil && sellerId == nil {
			rows, err = r.store.conn.Query(context.Background(), "SELECT * FROM offers AND name LIKE $1",
				*pattern + "%")
			if err != nil {
				return nil, err
			}
		} else if offerId == nil && sellerId != nil {
			rows, err = r.store.conn.Query(context.Background(), "SELECT * FROM offers WHERE seller_id = $1 AND name LIKE $2",
				sellerId, *pattern + "%")
			if err != nil {
				return nil, err
			}
		} else if offerId != nil && sellerId == nil {
			rows, err = r.store.conn.Query(context.Background(), "SELECT * FROM offers WHERE offer_id = $1 AND name LIKE $2",
				offerId, *pattern + "%")
			if err != nil {
				return nil, err
			}
		} else {
			rows, err = r.store.conn.Query(context.Background(), "SELECT * FROM offers WHERE offer_id = $1 AND seller_id = $2 AND name LIKE $3",
				offerId, sellerId, *pattern + "%")
			if err != nil {
				return nil, err
			}
		}
	}

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

func (r *OfferRepository) Delete(o *model.Offer, rs *model.RowsStats) error {
	_, err := r.store.conn.Exec(context.Background(), "DELETE FROM offers WHERE offer_id = $1", o.OfferId)
	rs.DeletedRows += 1
	return err
}

func (r *OfferRepository) Exists(offerId, sellerId int) bool {
	var exists bool
	_ = r.store.conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM offers WHERE offer_id = $1 AND seller_id = $2)",
		offerId, sellerId).Scan(&exists)
	fmt.Println(offerId, sellerId, exists)
	return exists
}

func (r *OfferRepository) Update(o *model.Offer, rs *model.RowsStats) error {
	_, err := r.store.conn.Exec(context.Background(), "UPDATE offers SET name = $1, price = $2, quantity = $3 WHERE seller_id = $4 and offer_id = $5",
		o.Name, o.Price, o.Quantity, o.SellerId, o.OfferId)
	rs.UpdatedRows += 1
	return err
}