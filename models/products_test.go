package models

import (
	"testing"
)

func TestProductValidation(t *testing.T) {
	p := Product{
		ID:          1,
		Name:        "Test Product",
		Description: "This is a test product",
		Price:       10.99,
		SKU:         "abc-xyz-def",
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Expected product to be valid, got error: %s", err)
	}

	// p2 := Product{}
	// err = p2.Validate()
	// if err == nil {
	// 	t.Error("Expected product to be invalid, got no error")
	// }
}
