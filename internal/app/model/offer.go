package model

type Offer struct {
	OfferId    int `json:"offer_id"`
	Name    string `json:"name"`
	Price      int `json:"price"`
	Quantity   int `json:"quantity"`
	Available bool `json:"-"`
	SellerId   int `json:"seller_id"`
}
