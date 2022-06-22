package controller

import (
	"net/http"

	"github.com/rizkyprawirap/Toko/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	productModel := models.Product{}
	products, err := productModel.GetProducts(server.DB)
	if err != nil {
		return // handle error
	}

	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products": products,
	})
}
