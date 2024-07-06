package entities

import (
	"reflect"
	"time"

	validate_object "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/validate_objects"
)

type PortfolioEntityParams struct {
	ID                  string
	Channel             string
	Country             string
	CreateAt            *time.Time
	CustomerID          string
	Route               string
	SKU                 string
	Title               string
	CategoryID          string
	Category            string
	Brand               string
	Classification      string
	UnitsPerBox         string
	MinOrderUnits       string
	PackageDescription  string
	PackageUnitDescription string
	QuantityMaxRedeem   int
	RedeemUnit          string
	OrderReasonRedeem   int
	SKURedeem           bool
	Price               float64
	Points              int
	Taxes               []*Tax
}

type Portfolio struct {
	id *validate_object.StringValidateObject
	channel *validate_object.StringValidateObject
	country *validate_object.StringValidateObject
	createAt *validate_object.TimeValidateObject
	customerId *validate_object.StringValidateObject
	route *validate_object.StringValidateObject
	sku *validate_object.StringValidateObject
	title *validate_object.StringValidateObject
	categoryId *validate_object.StringValidateObject
	category *validate_object.StringValidateObject
	brand *validate_object.StringValidateObject
	classification *validate_object.StringValidateObject
	unitsPerBox *validate_object.NumberValidateObject
	minOrderUnits *validate_object.FloatValidateObject
	packageDescription *validate_object.StringValidateObject
	packageUnitDescription *validate_object.StringValidateObject
	quantityMaxRedeem *validate_object.NumberValidateObject
	redeemUnit *validate_object.StringValidateObject
	orderReasonRedeem *validate_object.NumberValidateObject
	skuRedeem bool
	fullPrice *validate_object.FloatValidateObject
	points *validate_object.NumberValidateObject
	taxes []*Tax
}

func InitPortfolio(params *PortfolioEntityParams) *Portfolio {
	return &Portfolio{
		id: validate_object.NewStringValueObject(params.ID, "id").IsOptional().IsID(),
		channel: validate_object.NewStringValueObject(params.Channel, "channel").MinLength(4).MaxLength(10),
		country: validate_object.NewStringValueObject(params.Country, "country").MinLength(4).MaxLength(5),
		createAt: validate_object.NewTimeValueObject(params.CreateAt, "createAt").Format("2006-01-02T00:00:00-07:00"),
		customerId: validate_object.NewStringValueObject(params.CustomerID, "customerId").IsID(),
		route: validate_object.NewStringValueObject(params.Route, "route").MinLength(4).MaxLength(10),
		sku: validate_object.NewStringValueObject(params.SKU, "sku").MinLength(6).MaxLength(10),
		title: validate_object.NewStringValueObject(params.Title, "title").MinLength(4).MaxLength(50).TransformUpperCase(),
		categoryId: validate_object.NewStringValueObject(params.CategoryID, "categoryId").TransformSnakeCase(),
		category: validate_object.NewStringValueObject(params.Category, "category").MinLength(4).MaxLength(70).TransformUpperCase(),
		brand: validate_object.NewStringValueObject(params.Brand, "brand").MinLength(4).MaxLength(70).TransformLowerCase(),
		classification: validate_object.NewStringValueObject(params.Classification, "classification").MinLength(4).MaxLength(70).TransformUpperCase(),
		unitsPerBox: validate_object.NewNumberValueObject(params.UnitsPerBox, "unitsPerBox").IsPositive(),
		minOrderUnits: validate_object.NewFloatValueObject(params.MinOrderUnits, "minOrderUnits").IsPositive().Decimals(2),
		packageDescription: validate_object.NewStringValueObject(params.PackageDescription, "packageDescription").MinLength(4).MaxLength(70).TransformUpperCase(),
		packageUnitDescription: validate_object.NewStringValueObject(params.PackageUnitDescription, "packageUnitDescription").MinLength(4).MaxLength(70).TransformUpperCase(),
		quantityMaxRedeem: validate_object.NewNumberValueObject(params.QuantityMaxRedeem, "quantityMaxRedeem").IsPositive(),
		redeemUnit: validate_object.NewStringValueObject(params.RedeemUnit, "redeemUnit").MinLength(4).MaxLength(10),
		orderReasonRedeem: validate_object.NewNumberValueObject(params.OrderReasonRedeem, "orderReasonRedeem").IsPositive(),
		skuRedeem: params.SKURedeem,
		fullPrice: validate_object.NewFloatValueObject(params.Price, "fullPrice").IsPositive().Decimals(2),
		points: validate_object.NewNumberValueObject(params.Points, "points").IsPositive(),
		taxes: params.Taxes,
	}
}

func (p *Portfolio) ID() string {
	return p.id.Value()
}

func (p *Portfolio) Channel() string {
	return p.channel.Value()
}

func (p *Portfolio) Country() string {
	return p.country.Value()
}

func (p *Portfolio) CreateAt() time.Time {
	return p.createAt.Value()
}

func (p *Portfolio) CustomerID() string {
	return p.customerId.Value()
}

func (p *Portfolio) Route() string {
	return p.route.Value()
}

func (p *Portfolio) SKU() string {
	return p.sku.Value()
}

func (p *Portfolio) Title() string {
	return p.title.Value()
}

func (p *Portfolio) CategoryID() string {
	return p.categoryId.Value()
}

func (p *Portfolio) Category() string {
	return p.category.Value()
}

func (p *Portfolio) Brand() string {
	return p.brand.Value()
}

func (p *Portfolio) Classification() string {
	return p.classification.Value()
}

func (p *Portfolio) UnitsPerBox() int {
	return p.unitsPerBox.Value()
}

func (p *Portfolio) MinOrderUnits() float64 {
	return p.minOrderUnits.Value()
}

func (p *Portfolio) PackageDescription() string {
	return p.packageDescription.Value()
}

func (p *Portfolio) PackageUnitDescription() string {
	return p.packageUnitDescription.Value()
}

func (p *Portfolio) QuantityMaxRedeem() int {
	return p.quantityMaxRedeem.Value()
}

func (p *Portfolio) RedeemUnit() string {
	return p.redeemUnit.Value()
}

func (p *Portfolio) OrderReasonRedeem() int {
	return p.orderReasonRedeem.Value()
}

func (p *Portfolio) SKURedeem() bool {
	return p.skuRedeem
}

func (p *Portfolio) Price() float64 {
	value:=p.fullPrice.Value()

	if (p.taxes!=nil) {
		for i := 0; i < len(p.taxes); i++ {
			value = value	* (1 + p.taxes[i].Rate())
		}
	}
	
	return value
}

func (p *Portfolio) Points() int {
	return p.points.Value()
}

func (p *Portfolio) Validate() error {
	v := reflect.ValueOf(p)
	
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)
	
		if field.Name != "taxes" {
			if validator, ok := value.Interface().(interface{ Validate() error }); ok {
				if err := validator.Validate(); err != nil {
					return err
				}
			}
		}

		if field.Name == "taxes"{
			if p.taxes == nil {
				return nil
			}

			for j := 0; j < len(p.taxes); j++ {
				if err := p.taxes[j].Validate(); err != nil {
					return err
				}
			}
		}
	}
	
	return nil
}