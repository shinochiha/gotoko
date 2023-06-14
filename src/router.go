package src

import (
	"github.com/shinochiha/gotoko/app"
	"github.com/shinochiha/gotoko/src/contact"
	"github.com/shinochiha/gotoko/src/product"
	// import : DONT REMOVE THIS COMMENT
)

func Router() *routerUtil {
	if router == nil {
		router = &routerUtil{}
		router.Configure()
		router.isConfigured = true
	}
	return router
}

var router *routerUtil

type routerUtil struct {
	isConfigured bool
}

func (r *routerUtil) Configure() {
	app.Server().AddRoute("/api/version", "GET", app.VersionHandler, nil)

	app.Server().AddRoute("/api/v1/contacts", "POST", contact.REST().Create, contact.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/contacts", "GET", contact.REST().Get, contact.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/contacts/{id}", "GET", contact.REST().GetByID, contact.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/contacts/{id}", "PUT", contact.REST().UpdateByID, contact.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/contacts/{id}", "PATCH", contact.REST().PartiallyUpdateByID, contact.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/contacts/{id}", "DELETE", contact.REST().DeleteByID, contact.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/products", "POST", product.REST().Create, product.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/products", "GET", product.REST().Get, product.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/products/{id}", "GET", product.REST().GetByID, product.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/products/{id}", "PUT", product.REST().UpdateByID, product.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/products/{id}", "PATCH", product.REST().PartiallyUpdateByID, product.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/products/{id}", "DELETE", product.REST().DeleteByID, product.OpenAPI().DeleteByID())

	// AddRoute : DONT REMOVE THIS COMMENT
}
