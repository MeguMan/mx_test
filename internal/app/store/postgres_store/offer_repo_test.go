package postgres_store

import (
	"github.com/MeguMan/mx_test/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOfferRepository_Create(t *testing.T) {
	conn, teardown := TestDB(t)
	defer teardown("offers")
	s := New(conn)
	o := model.Offer{
		OfferId:   555,
		Name:      "TestingName",
		Price:     555,
		Quantity:  555,
		Available: true,
		SellerId:  666,
	}
	rs := model.RowsStats{}
	err := s.Offer().Create(&o, &rs)
	assert.NoError(t, err)
	assert.Equal(t, 1, rs.CreatedRows)
}

func TestOfferRepository_Delete(t *testing.T) {
	conn, teardown := TestDB(t)
	defer teardown("offers")
	s := New(conn)
	o := model.Offer{
		OfferId:   555,
		Name:      "TestingName",
		Price:     555,
		Quantity:  555,
		Available: true,
		SellerId:  666,
	}
	rs := model.RowsStats{}
	_ = s.Offer().Create(&o, &rs)
	err := s.Offer().Delete(&o, &rs)
	assert.NoError(t, err)
	assert.Equal(t, 1, rs.DeletedRows)
}

func TestOfferRepository_Exists(t *testing.T) {
	conn, teardown := TestDB(t)
	defer teardown("offers")
	s := New(conn)
	o := model.Offer{
		OfferId:   555,
		Name:      "TestingName",
		Price:     555,
		Quantity:  555,
		Available: true,
		SellerId:  666,
	}
	rs := model.RowsStats{}
	_ = s.Offer().Create(&o, &rs)
	exists := s.Offer().Exists(o.OfferId, o.SellerId)
	assert.True(t, exists)
	_ = s.Offer().Delete(&o, &rs)
	exists = s.Offer().Exists(o.OfferId, o.SellerId)
	assert.False(t, exists)
}

func TestOfferRepository_GetByPattern(t *testing.T) {
	conn, teardown := TestDB(t)
	defer teardown("offers")
	s := New(conn)
	o := model.Offer{
		OfferId:   555,
		Name:      "TestingName",
		Price:     555,
		Quantity:  555,
		Available: true,
		SellerId:  666,
	}
	rs := model.RowsStats{}
	_ = s.Offer().Create(&o, &rs)
	pattern := ""
	oo, err := s.Offer().GetByPattern(&o.OfferId, &o.SellerId, &pattern)
	assert.NoError(t, err)
	assert.NotNil(t, oo)
	_ = s.Offer().Delete(&o, &rs)
	oo, err = s.Offer().GetByPattern(&o.OfferId, &o.SellerId, &pattern)
	assert.NoError(t, err)
	assert.Nil(t, oo)
}

func TestOfferRepository_Update(t *testing.T) {
	conn, teardown := TestDB(t)
	defer teardown("offers")
	s := New(conn)
	o := model.Offer{
		OfferId:   555,
		Name:      "TestingName",
		Price:     555,
		Quantity:  555,
		Available: true,
		SellerId:  666,
	}
	rs := model.RowsStats{}
	_ = s.Offer().Create(&o, &rs)
	o.Name = "newTestingName"
	err := s.Offer().Update(&o, &rs)
	assert.NoError(t, err)
	assert.Equal(t, 1, rs.UpdatedRows)
}
