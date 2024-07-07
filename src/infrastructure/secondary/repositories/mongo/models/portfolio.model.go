package models_mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Portfolio struct {
	ID          primitive.ObjectID     `bson:"_id"`
	Channel     string     `bson:"channel"`
	Country     string     `bson:"country"`
	CreatedDate *time.Time `bson:"createdDate,omitempty"`
	CustomerCode string   `bson:"customerCode"`
	Route       string     `bson:"route"`
	Sku         string     `bson:"sku"`
	Title       string     `bson:"title"`
	CategoryId  string     `bson:"categoryId"`
	Category    string     `bson:"category"`
	Brand       string     `bson:"brand"`
	Classification string  `bson:"classification"`
	UnitsPerBox int       `bson:"unitsPerBox"`
	MinOrderUnits float64 `bson:"minOrderUnits"`
	PackageDescription string `bson:"packageDescription"`
	PackageUnitDescription string `bson:"packageUnitDescription"`
	QuantityMaxRedeem int `bson:"quantityMaxRedeem"`
	RedeemUnit string `bson:"redeemUnit"`
	OrderReasonRedeem int `bson:"orderReasonRedeem"`
	SkuRedeem bool `bson:"skuRedeem"`
	Price Price `bson:"price"`
	Points int `bson:"points"`
}

type Price struct {
	FullPrice float64 `bson:"fullPrice"`
	Taxes     []Tax `bson:"taxes"`
}

type Tax struct {
	TaxType string `bson:"taxType"`
	TaxId   string `bson:"taxId"`
	Rate    int `bson:"rate"`
}

type Indexes struct {
	IndexName string `bson:"indexName"`
	Unique    bool   `bson:"unique"`
	Keys      []string `bson:"keys"`
}