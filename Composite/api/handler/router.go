package handler

import (
	handlerIn "product_service_composite/api/handler/httpCalls/product"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// MakeProductHandler  handler for product routes
func MakeProductHandler(r *mux.Router, n negroni.Negroni, productHandlerOut *handlerIn.HandlerProduct) {
	r.Handle("/product", n.With(
		negroni.Wrap(productHandlerOut.CreateProduct()),
	)).Methods("POST", "OPTIONS").Name("CreateProduct")

	r.Handle("/product/{sku}", n.With(
		negroni.Wrap(productHandlerOut.DeleteProduct()),
	)).Methods("DELETE", "OPTIONS").Name("DeleteProduct")

	r.Handle("/product/{sku}", n.With(
		negroni.Wrap(productHandlerOut.UpdateProduct()),
	)).Methods("PATCH", "OPTIONS").Name("UpdateProduct")

	r.Handle("/product/{sku}", n.With(
		negroni.Wrap(productHandlerOut.GetProduct()),
	)).Methods("GET", "OPTIONS").Name("GetProduct")

	r.Handle("/product", n.With(
		negroni.Wrap(productHandlerOut.GetAllProduct()),
	)).Methods("GET", "OPTIONS").Name("GetAllProduct")
}
