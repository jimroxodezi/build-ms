// Package classification of Product API
//
// Documentation for product API.
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jimroxodezi/build-ms/internal/models"
	"github.com/jimroxodezi/build-ms/internal/service"
	"github.com/go-chi/chi/v5"
)

// ProductHandler is the handler for products
// we use a service layer to handle the business logic
type ProductHandler struct {
	s service.ProductService
	l *log.Logger
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{l: l}
}

// swagger:route GET /products products listProducts
// Returns a list of products.
//
// responses:
//   200: productsResponse}
// GetProducts returns the list of products
func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	// p.l.Println("Handle GET products")
	w.Header().Set("Content-Type", "application/json")
	productList, _ := p.s.GetProducts(ctx)
	err := json.NewEncoder(w).Encode(productList)
	if err != nil {
		// http.Error(w, "Error marshaling products", http.StatusInternalServerError)
		WriteError(w, http.StatusInternalServerError, "Error marshaling products")
		return
	}
}

// GetProduct returns a single product by ID
func (p *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	product, err := p.s.GetProduct(ctx, id)
	if err != nil {
		WriteError(w, http.StatusNotFound, "Product not found")
		return
	}
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error marshaling product")
		return
	}
}



// swagger:route POST /products products addProduct
// Adds a new product to the list.
//
// responses:
//   201: noContentResponse
// AddProduct adds a new product to the list
func (p *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		// http.Error(w, "Error unmarshaling product", http.StatusBadRequest)
		WriteError(w, http.StatusBadRequest, "Error unmarshaling product")
		return
	}

	err = product.Validate()
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Error validating product")
		return
	}

	err = p.s.CreateProduct(ctx, product)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Error creating product")
		return
	}
}

// swagger:route PUT /products/{id} products updateProduct
// Updates an existing product in the list.
//
// responses:
//   204: noContentResponse
// UpdateProducts updates an existing product in the list
func (p *ProductHandler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// http.Error(w, "Error parsing product id", http.StatusBadRequest)
		WriteError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	
	product := models.Product{}
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Error unmarshaling product")
		return
	}

	err = product.Validate()
	if err != nil {
		http.Error(w, "Error validating product", http.StatusBadRequest)
		return
	}

	err = p.s.UpdateProduct(ctx, id, product)
	if err == models.ErrProductNotFound {
		http.Error(w, "Error updating product", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}	
}


// // contextKey is a custom type to avoid context key collisions
// type contextKey struct{}
// // productKey is the key used to store the product in the context
// var productKey = contextKey{}

// // ProductValidate is a middleware that validates the product in the request body and stores it in the context
// func (p *Products) ProductValidate(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		product := models.Product{}
// 		err := product.FromJSON(r.Body)
// 		if err != nil {
// 			http.Error(w, "Error unmarshaling product", http.StatusBadRequest)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), productKey, product)
// 		r = r.WithContext(ctx)

// 		next.ServeHTTP(w, r)
// 	})
// }
