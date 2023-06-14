package src

import (
	"github.com/shinochiha/gotoko/app"
	"github.com/shinochiha/gotoko/src/contact"
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

	// AddRoute : DONT REMOVE THIS COMMENT
}
