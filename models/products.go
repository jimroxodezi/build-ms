package models

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	
	// "github.com/build-ms/service"
)

var (
	// validate is the validator instance used for validating products
	validate = validator.New()

	// ErrProductNotFound is the error returned when a product is not found
	ErrProductNotFound = fmt.Errorf("Product not found")

	// skuRegexp is the regular expression used to validate product SKUs
	skuRegexp = regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
)


func Init() {
	err := validate.RegisterValidation("sku", validateSKU)
	if err != nil {
		panic(err)
	}
}

// swagger:model Product
// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"required,gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   time.Time  `json:"-"`
	UpdatedOn   time.Time  `json:"-"`
	DeletedOn   time.Time  `json:"-"`
}

func (p *Product) Validate() error {
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	return skuRegexp.MatchString(fl.Field().String())
}


// Products is a slice of Product pointers
type Products []*Product
